package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"

	"github.com/fiorix/wsdl2go/soap"
	"github.com/regiohelden/innovazammad/innovaphone"
	"github.com/sirupsen/logrus"
)

type innovaphoneSession struct {
	innovaphone.PbxPortType
	sessionHandle     int
	flipCallDirection bool
}

func newInnovaphoneService() innovaphoneSession {
	client := &soap.Client{
		URL:       fmt.Sprintf("%s/%s", strings.TrimRight(conf.PBX.URL, "/"), strings.TrimLeft(conf.PBX.EndpointPath, "/")),
		Namespace: innovaphone.Namespace,
		Pre: func(r *http.Request) {
			r.SetBasicAuth(conf.PBX.User, conf.PBX.Pass)
		},
	}
	logrus.Infof("initializing service for URL: %s", client.URL)
	return innovaphoneSession{
		PbxPortType: innovaphone.NewPbxPortType(client),
	}
}

func (service *innovaphoneSession) tryConnectionInit() {
	if err := backoff.RetryNotify(service.connectionInit, backoff.NewExponentialBackOff(), func(err error, dur time.Duration) {
		logrus.Errorf("%s (will retry in %s)", err, dur)
	}); err != nil {
		// we retry forever, so we should never hit this
		panic("unexpected error while retrying")
	}
}

func (service *innovaphoneSession) connectionInit() error {
	_, gatekeeperID, location, version, serial, err := service.Version()
	if err != nil {
		return errors.Wrap(err, "could not connect to PBX")
	}
	logrus.Infof("connected to %s, %s, version '%s', serial '%s'\n", gatekeeperID, location, strings.TrimSpace(version), serial)

	sessionHandle, key, err := service.Initialize(conf.PBX.MonitorUser, conf.PBX.AppName, true, true, true, true, true)
	if err != nil {
		return errors.Wrap(err, "could not initialize PBX session")
	}
	service.sessionHandle = sessionHandle

	if ret, err := service.Echo(sessionHandle, key); err != nil || ret == 0 {
		return errors.Wrapf(err, "could not verify session initialization via echo (code %d)", ret)
	}

	if _, err = service.UserInitialize(sessionHandle, conf.PBX.MonitorUser, false, true, ""); err != nil {
		return errors.Wrap(err, "could not initialize PBX user session")
	}

	// we have to poll once to get the current UserInfo, which in turn is necessary to know if we're monitoring a normal
	// or gateway user (i.e.: if we have to flip "this" and "peer" numbers)
	pollResp, err := service.Poll(service.sessionHandle)
	if err != nil {
		return errors.Wrap(err, "error while polling")
	}
	if len(pollResp.User.Items) < 1 {
		return errors.New("first poll did not return any users")
	}
	// first user is the one we've initialized above
	switch pollResp.User.Items[0].Type {
	case "ep":
		// docs: "Note that the PBX API is PBX-centric, not terminal centric. As such, it considers a call from the PBX to
		// 				the terminal as outbound."
		service.flipCallDirection = true
	}

	return nil
}

func (service *innovaphoneSession) pollForever() error {
	logrus.Info("polling for call events")
	for {
		pollResp, err := service.Poll(service.sessionHandle)
		if err != nil {
			return errors.Wrap(err, "error while polling")
		}
		for _, call := range pollResp.Call.Items {
			go service.handleCall(call)
		}
	}
}

func (service *innovaphoneSession) handleCall(call *innovaphone.CallInfo) {
	cis := &callInSession{
		innovaphoneSession: service,
		CallInfo:           call,
	}
	src, dst := cis.GetSourceAndDestination()
	dir := cis.GetDirection()

	logrus.Debugf("%s handling %s → %s (%s/%s/%s)", call, src, dst, dir, cis.GetState(), call.GetHold())

	local := src
	if dir == innovaphone.DirectionInbound {
		local = dst
	}
	if !service.shouldHandleNumber(local) {
		logrus.Debugf("skipping call for number not in group: '%s' not in '%s'", dst.E164, conf.PBX.FilterOnGroup)
		return
	}

	submitToZammad(cis)
}

type callInSession struct {
	*innovaphoneSession
	*innovaphone.CallInfo
}

// this is here instead of innovaphone package because direction is relative to the monitor user
func (call *callInSession) GetDirection() innovaphone.Direction {
	dir := call.CallInfo.GetDirection()
	if call.flipCallDirection {
		dir = dir.Flip()
	}
	return dir
}

// used for caching FindUser results
var userCache sync.Map

type userCacheEntry struct {
	groups    []*innovaphone.Group
	timestamp time.Time
}

// shouldHandleNumber decides whether a call for a
func (service *innovaphoneSession) shouldHandleNumber(no *innovaphone.No) bool {
	if conf.PBX.FilterOnGroup != "" {
		if no.E164 == "" {
			return false
		}

		var groups []*innovaphone.Group
		var cacheMiss = true
		if time.Duration(*conf.PBX.GroupCacheTime) > time.Duration(0) {
			// try to load groups from cache
			if val, ok := userCache.Load(no.E164); ok {
				cacheEntry, ok := val.(userCacheEntry)
				if !ok {
					panic("retrieved unknown value from userCache")
				}
				if cacheEntry.timestamp.After(time.Now().Add(-time.Duration(*conf.PBX.GroupCacheTime))) {
					groups = cacheEntry.groups
					cacheMiss = false
				}
			}
		}

		if cacheMiss {
			logrus.Debugf("searching for number %s", no.E164)
			userArray, err := service.FindUser("", "", "", "", no.Cn, "", no.E164, 1, 0)
			if err != nil {
				logrus.Errorf("error finding number '%s': %s", no.E164, err)
				return false
			}
			if len(userArray.Items) < 1 {
				logrus.Warnf("could not find number '%s'; skipping call", no.E164)
				return false
			}
			groups = userArray.Items[0].Groups.Items
			userCache.Store(no.E164, userCacheEntry{
				groups:    groups,
				timestamp: time.Now(),
			})
		}

		for _, group := range groups {
			if group.Group == conf.PBX.FilterOnGroup {
				return true
			}
		}
		return false
	}
	return true
}

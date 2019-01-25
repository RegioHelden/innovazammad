package innovaphone

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
	"github.com/regiohelden/innovazammad/config"

	"github.com/fiorix/wsdl2go/soap"
	"github.com/sirupsen/logrus"
)

// Session keeps information about the current session to the innovaphone PBX
type Session struct {
	service           PbxPortType
	sessionHandle     int
	flipCallDirection bool
}

// NewSession returns a new innovaphone session with context information.
// It will retry indefinitely to connect and confirm the connection to the PBX.
func NewSession() Session {
	client := &soap.Client{
		URL:       fmt.Sprintf("%s/%s", strings.TrimRight(config.Global.PBX.URL, "/"), strings.TrimLeft(config.Global.PBX.EndpointPath, "/")),
		Namespace: Namespace,
		Pre: func(r *http.Request) {
			r.SetBasicAuth(config.Global.PBX.User, config.Global.PBX.Pass)
		},
	}
	logrus.Infof("initializing session for URL: %s", client.URL)
	is := Session{
		service: NewPbxPortType(client),
	}

	if err := backoff.RetryNotify(is.connectionInit, backoff.NewExponentialBackOff(), func(err error, dur time.Duration) {
		logrus.Errorf("%s (will retry in %s)", err, dur)
	}); err != nil {
		// we retry forever, so we should never hit this
		panic("unexpected error while retrying")
	}

	return is
}

func (session *Session) connectionInit() error {
	_, gatekeeperID, location, version, serial, err := session.service.Version()
	if err != nil {
		return errors.Wrap(err, "could not connect to PBX")
	}
	logrus.WithFields(logrus.Fields{
		"id":       gatekeeperID,
		"location": location,
		"version":  strings.TrimSpace(version),
		"serial":   serial,
	}).Info("connection established")

	sessionHandle, key, err := session.service.Initialize(config.Global.PBX.MonitorUser, config.Global.PBX.AppName, true, true, true, true, true)
	if err != nil {
		return errors.Wrap(err, "could not initialize PBX session")
	}
	session.sessionHandle = sessionHandle

	if ret, err := session.service.Echo(sessionHandle, key); err != nil || ret == 0 {
		return errors.Wrapf(err, "could not verify session initialization via echo (code %d)", ret)
	}

	if _, err = session.service.UserInitialize(sessionHandle, config.Global.PBX.MonitorUser, false, true, ""); err != nil {
		return errors.Wrap(err, "could not initialize PBX user session")
	}

	// we have to poll once to get the current UserInfo, which in turn is necessary to know if we're monitoring a normal
	// or gateway user (i.e.: if we have to flip "this" and "peer" numbers)
	pollResp, err := session.service.Poll(session.sessionHandle)
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
		session.flipCallDirection = true
	}

	return nil
}

// PollForever will return one CallInSession per successful poll.
// If it encounters an error, it will return it over the errors channel and cease polling.
func (session *Session) PollForever() (<-chan *CallInSession, <-chan error) {
	logrus.Info("polling for call events")
	errs := make(chan error)
	calls := make(chan *CallInSession)

	go func() {
		for {
			pollResp, err := session.service.Poll(session.sessionHandle)
			if err != nil {
				errs <- errors.Wrap(err, "error while polling")
				break
			}
			for _, call := range pollResp.Call.Items {
				cis := &CallInSession{
					Session:  session,
					CallInfo: call,
				}
				src, dst := cis.GetSourceAndDestination()
				dir := cis.GetDirection()

				logrus.WithFields(logrus.Fields{
					"call":      call,
					"source":    src,
					"dest":      dst,
					"direction": dir,
					"state":     cis.GetState(),
					"hold":      cis.GetHold(),
				}).Debug("received event")

				calls <- cis
			}
		}
	}()

	return calls, errs
}

// CallInSession puts call information together with the session context. Some decisions about calls can only be made
// in context, since - for instance - the direction information is dependent on the user being monitored.
type CallInSession struct {
	*Session
	*CallInfo
}

// GetDirection retuns the call direction as interpreted by an outside observer. It wraps the PBX' notion of direction,
// which might be relative to itself.
func (call *CallInSession) GetDirection() Direction {
	dir := call.CallInfo.GetDirection()
	if call.flipCallDirection {
		dir = dir.Flip()
	}
	return dir
}

// used for caching FindUser results
var userCache sync.Map

type userCacheEntry struct {
	groups    []*Group
	timestamp time.Time
}

// ShouldHandle decides whether a call involves any of the groups being filtered on (see Config.FilterOnGroup)
func (call *CallInSession) ShouldHandle() bool {
	if config.Global.PBX.FilterOnGroup != "" {
		src, dst := call.GetSourceAndDestination()
		no := src
		if call.GetDirection() == DirectionInbound {
			no = dst
		}
		if no.E164 == "" {
			return false
		}

		var groups []*Group
		var cacheMiss = true
		if time.Duration(*config.Global.PBX.GroupCacheTime) > time.Duration(0) {
			// try to load groups from cache
			if val, ok := userCache.Load(no.E164); ok {
				cacheEntry, ok := val.(userCacheEntry)
				if !ok {
					panic("retrieved unknown value from userCache")
				}
				if cacheEntry.timestamp.After(time.Now().Add(-time.Duration(*config.Global.PBX.GroupCacheTime))) {
					groups = cacheEntry.groups
					cacheMiss = false
				}
			}
		}

		if cacheMiss {
			logrus.WithField("call", call).Debugf("searching for number %s", no.E164)
			userArray, err := call.service.FindUser("", "", "", "", no.Cn, "", no.E164, 1, 0)
			if err != nil {
				logrus.WithField("call", call).Errorf("error finding number '%s': %s", no.E164, err)
				return false
			}
			if len(userArray.Items) < 1 {
				logrus.WithField("call", call).Warnf("could not find number '%s'; skipping call", no.E164)
				return false
			}
			groups = userArray.Items[0].Groups.Items
			userCache.Store(no.E164, userCacheEntry{
				groups:    groups,
				timestamp: time.Now(),
			})
		}

		for _, group := range groups {
			if group.Group == config.Global.PBX.FilterOnGroup {
				return true
			}
		}
		return false
	}
	return true
}

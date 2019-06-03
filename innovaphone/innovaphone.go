package innovaphone

import (
	"context"
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

//go:generate wsdl2go -p innovaphone -i pbx.wsdl -o pbx.generated.go

// Session keeps information about the current session to the innovaphone PBX
type Session struct {
	PbxPortType
	sessionHandle     int
	flipCallDirection bool
}

// NewSession returns a new innovaphone session with context information.
// It will retry indefinitely to connect and confirm the connection to the PBX.
func NewSession(ctx context.Context) *Session {
	client := &soap.Client{
		URL:       fmt.Sprintf("%s/%s", strings.TrimRight(config.Global.PBX.URL, "/"), strings.TrimLeft(config.Global.PBX.EndpointPath, "/")),
		Namespace: Namespace,
		Pre: func(r *http.Request) {
			r.SetBasicAuth(config.Global.PBX.User, config.Global.PBX.Pass)
			*r = *r.WithContext(ctx)
		},
	}
	logrus.Infof("initializing session for URL: %s", client.URL)
	is := &Session{
		PbxPortType: NewPbxPortType(client),
	}

	if err := backoff.RetryNotify(is.connectionInit, backoff.WithContext(backoff.NewExponentialBackOff(), ctx), func(err error, dur time.Duration) {
		logrus.Errorf("%s (will retry in %s)", err, dur)
	}); err != nil {
		return nil
	}

	return is
}

func (session *Session) connectionInit() error {
	versionResp, err := session.Version(&Version{})
	if err != nil {
		return errors.Wrap(err, "could not connect to PBX")
	}
	logrus.WithFields(logrus.Fields{
		"id":       versionResp.GatekeeperID,
		"location": versionResp.Location,
		"version":  strings.TrimSpace(*versionResp.FirmwareVersion),
		"serial":   versionResp.SerialNumber,
	}).Info("connection established")

	initializeResp, err := session.Initialize(&Initialize{
		User:   &config.Global.PBX.MonitorUser,
		Appl:   &config.Global.PBX.AppName,
		V:      true,
		V501:   true,
		V700:   true,
		V800:   true,
		Vx1100: true,
	})
	if err != nil {
		return errors.Wrap(err, "could not initialize PBX session")
	}
	session.sessionHandle = initializeResp.Return

	if echoResp, err := session.Echo(&Echo{
		Session: initializeResp.Return,
		Key:     initializeResp.Key,
	}); err != nil || echoResp.Return == 0 {
		return errors.Wrapf(err, "could not verify session initialization via echo (code %d)", echoResp.Return)
	}

	if _, err = session.UserInitialize(&UserInitialize{
		Session: session.sessionHandle,
		User:    &config.Global.PBX.MonitorUser,
		Xfer:    false,
		Disc:    true,
	}); err != nil {
		return errors.Wrap(err, "could not initialize PBX user session")
	}

	// we have to poll once to get the current UserInfo, which in turn is necessary to know if we're monitoring a normal
	// or gateway user (i.e.: if we have to flip "this" and "peer" numbers)
	pollResp, err := session.Poll(&Poll{Session: session.sessionHandle})
	if err != nil {
		return errors.Wrap(err, "error while polling")
	}
	if len(pollResp.Return.User) < 1 {
		return errors.New("first poll did not return any users")
	}
	// first user is the one we've initialized above
	switch pollResp.Return.User[0].Type {
	case "ep":
		// docs: "Note that the PBX API is PBX-centric, not terminal centric. As such, it considers a call from the PBX to
		// 				the terminal as outbound."
		session.flipCallDirection = true
	}

	return nil
}

// IsDirectionFlipped returns whether the current session should treat directions as inverted
func (session *Session) IsDirectionFlipped() bool {
	return session.flipCallDirection
}

// PollForever will return one CallInSession per successful poll.
// If it encounters an error, it will return it over the errors channel and cease polling.
func (session *Session) PollForever() (<-chan *CallInSession, <-chan error) {
	if session == nil {
		return nil, nil
	}

	logrus.Infof("polling for call events for user %s", config.Global.PBX.MonitorUser)
	errs := make(chan error)
	calls := make(chan *CallInSession)

	go func() {
		for {
			pollResp, err := session.Poll(&Poll{Session: session.sessionHandle})
			if err != nil {
				errs <- errors.Wrap(err, "error while polling")
				break
			}
			for _, call := range pollResp.Return.Call {
				cis := &CallInSession{
					sessionInterface: session,
					CallInfo:         call,
				}
				src, dst, err := cis.GetSourceAndDestination()
				if err != nil {
					logrus.WithField("call", call).Warn(err)
					continue
				}
				dir := cis.GetDirection()

				logrus.WithFields(logrus.Fields{
					"call":      call,
					"source":    src.Normalize(),
					"dest":      dst.Normalize(),
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
	sessionInterface
	*CallInfo
}

type sessionInterface interface {
	IsDirectionFlipped() bool
	FindUser(*FindUser) (*FindUserResponse, error)
}

// GetDirection returns the call direction as interpreted by an outside observer. It wraps the PBX' notion of direction,
// which might be relative to itself.
func (call *CallInSession) GetDirection() Direction {
	dir := call.CallInfo.GetDirection()
	if call.IsDirectionFlipped() {
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
		src, dst, err := call.GetSourceAndDestination()
		if err != nil {
			logrus.WithField("call", call).Warn(err)
			return false
		}
		no := src
		if call.GetDirection() == DirectionInbound {
			no = dst
		}
		if no.Cn == "" {
			return false
		}

		var groups []*Group
		var cacheMiss = true
		if time.Duration(*config.Global.PBX.GroupCacheTime) > time.Duration(0) {
			// try to load groups from cache
			if val, ok := userCache.Load(no.Cn); ok {
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
			logrus.WithField("call", call).Debugf("searching for PBX object '%s'", no.Cn)
			nonEmpty := "x" // the API is weird ¯\_(ツ)_/¯
			findUserResp, err := call.FindUser(&FindUser{
				V501:   &nonEmpty,
				V700:   &nonEmpty,
				V800:   &nonEmpty,
				Vx1100: &nonEmpty,
				Cn:     &no.Cn,
				Count:  1,
				Nohide: true,
			})
			if err != nil {
				logrus.WithField("call", call).Errorf("error finding PBX object '%s': %s", no.Cn, err)
				return false
			}
			if len(findUserResp.Return.User) < 1 || findUserResp.Return.User[0].Cn != no.Cn {
				logrus.WithField("call", call).Warnf("could not find PBX object '%s'; skipping call", no.Cn)
				return false
			}
			groups = findUserResp.Return.User[0].Groups
			userCache.Store(no.Cn, userCacheEntry{
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

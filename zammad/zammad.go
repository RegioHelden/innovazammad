package zammad

import (
	"context"
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/regiohelden/innovazammad/config"
	"github.com/regiohelden/innovazammad/innovaphone"
	"github.com/sirupsen/logrus"
)

// State represents the call state from Zammad's perspective
type State int

// Possible states for a Zammad call
const (
	StateNew State = iota
	StateRinging
	StateConnected
	StateDisconnected
)

func (s State) String() string {
	switch s {
	case StateNew:
		return "StateNew"
	case StateRinging:
		return "StateRinging"
	case StateConnected:
		return "StateConnected"
	case StateDisconnected:
		return "StateDisconnected"
	default:
		panic("what!?")
	}
}

// MarshalJSON is used by expvar
func (s State) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", s.String())), nil
}

type callEntry struct {
	State State
	UUID  uuid.UUID
	sync.Mutex
}

type callMap map[int]*callEntry

// String is used by expvar
func (cs callMap) String() string {
	out, err := json.Marshal(cs)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

// Session keeps track of the states successfully submitted to Zammad
type Session struct {
	callMap            callMap
	stats              *expvar.Map
	callHandlingCancel context.CancelFunc
	sync.Mutex
}

type stateTransition struct {
	curState State
	newState innovaphone.State
}

// NewSession initializes a new Session to keep track of call information sent to zammad.
// The result should be used as a singleton.
func NewSession() (*Session, context.Context) {
	callHandlingCtx, callHandlingCancel := context.WithCancel(context.Background())
	calls := &Session{
		callMap:            callMap{},
		stats:              expvar.NewMap("stats"),
		callHandlingCancel: callHandlingCancel,
	}
	expvar.Publish("calls", calls.callMap)
	return calls, callHandlingCtx
}

// Get returns a State for the given callID. If no state is known, will return StateNew.
func (zs *Session) get(callID int) (*callEntry, error) {
	zs.Lock()
	defer zs.Unlock()
	if _, ok := zs.callMap[callID]; !ok {
		UUID, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}
		return &callEntry{
			State: StateNew,
			UUID:  UUID,
		}, nil
	}
	return zs.callMap[callID], nil
}

func (zs *Session) setOrUpdate(ctx context.Context, callID int, newEntry *callEntry) {
	if newEntry.State == StateDisconnected {
		// if we're shutting down and this was the last call, cancel the callHandlingCtx
		select {
		case <-ctx.Done():
			logrus.Debug("last call terminated; cancelling call handling")
			zs.callHandlingCancel()
		default:
		}
		zs.Lock()
		defer zs.Unlock()
		delete(zs.callMap, callID)
	} else if curEntry, ok := zs.callMap[callID]; ok {
		curEntry.State = newEntry.State
	} else {
		zs.stats.Add("calls_total", 1)
		zs.callMap[callID] = newEntry
	}
}

// ShutdownIfEmpty is used to short-circuit the call handling loop
func (zs *Session) ShutdownIfEmpty() {
	zs.Lock()
	defer zs.Unlock()
	if len(zs.callMap) == 0 {
		zs.callHandlingCancel()
	}
}

// Submit sends a call to Zammad, if we are aware of it and can correctly map its state to some known Zammad state
func (zs *Session) Submit(ctx context.Context, call *innovaphone.CallInSession) error {
	zs.stats.Add("events_total", 1)

	src, dst, user, err := call.GetSourceAndDestination()
	if err != nil {
		return err
	}
	dir := call.GetDirection()
	newState := call.GetState()

	entry, err := zs.get(call.Call)
	if err != nil {
		return err
	}
	// ensure we serialize access to one specific call
	entry.Lock()
	defer entry.Unlock()

	// only skip new calls, because a call-forwarding may send the call outside of the filtered group, meaning we have to
	// handle calls we would otherwise filter out based on involved numbers
	if entry.State == StateNew && !call.ShouldHandle() {
		logrus.WithField("call", call).Debugf("skipping call for number not in group '%s'", config.Global.PBX.FilterOnGroup)
		return nil
	}

	content := url.Values{
		"callId":    []string{entry.UUID.String()},
		"from":      []string{src.Normalize()},
		"to":        []string{dst.Normalize()},
		"direction": []string{dir.String()},
	}

	transition := stateTransition{curState: entry.State, newState: newState}
	switch transition {
	case
		stateTransition{StateNew, innovaphone.StateCallProc},
		stateTransition{StateNew, innovaphone.StateAlert}:
		// we do not get StateAlert on outgoing calls, so we have to assume StateCallProc is enough
		content.Set("event", "newCall")
		content.Set("user", user)
		entry.State = StateRinging
	case stateTransition{StateRinging, innovaphone.StateConnect}:
		content.Set("event", "answer")
		if dir == innovaphone.DirectionInbound {
			content.Set("user", user)
			content.Set("answeringNumber", dst.Normalize())
		}
		entry.State = StateConnected
	case stateTransition{StateRinging, innovaphone.StateDisconnectSent}:
		content.Set("event", "hangup")
		content.Set("cause", "cancel")
		entry.State = StateDisconnected
	case stateTransition{StateRinging, innovaphone.StateDisconnectReceived}:
		content.Set("event", "hangup")
		content.Set("cause", "noAnswer")
		entry.State = StateDisconnected
	case
		stateTransition{StateConnected, innovaphone.StateDisconnectSent},
		stateTransition{StateConnected, innovaphone.StateDisconnectReceived}:
		content.Set("event", "hangup")
		content.Set("cause", "normalClearing")
		entry.State = StateDisconnected
	default:
		logrus.WithField("call", call).Debugf("ignoring unknown state transition %s → %s", transition.curState, transition.newState)
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"call":    call,
		"content": content,
	}).Infof("submitting to Zammad")
	resp, err := http.PostForm(config.Global.Zammad.EndpointURL, content)
	if err != nil {
		return errors.Wrap(err, "could not submit to Zammad")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("could not submit to Zammad: %s", resp.Status)
	}
	zs.stats.Add("events_submitted", 1)
	zs.setOrUpdate(ctx, call.Call, entry)
	return nil
}

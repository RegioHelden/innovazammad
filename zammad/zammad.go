package zammad

import (
	"encoding/json"
	"expvar"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

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

type callStateMap map[int]State

// String is used by expvar
func (cs callStateMap) String() string {
	out, err := json.Marshal(cs)
	if err != nil {
		return err.Error()
	}
	return string(out)
}

// Session keeps track of the states successfully submitted to Zammad
type Session struct {
	callStateMap callStateMap
	sync.Mutex
}

type stateTransition struct {
	curState State
	newState innovaphone.State
}

// NewSession initializes a new Session to keep track of call information sent to zammad.
// The result should be used as a singleton.
func NewSession() *Session {
	calls := &Session{
		callStateMap: callStateMap{},
	}
	expvar.Publish("Session", calls.callStateMap)
	return calls
}

// Get returns a State for the given callID. If no state is known, will return StateNew.
func (zs *Session) get(callID int) State {
	if _, ok := zs.callStateMap[callID]; !ok {
		return StateNew
	}
	return zs.callStateMap[callID]
}

func (zs *Session) update(callID int, state State) {
	if state == StateDisconnected {
		delete(zs.callStateMap, callID)
	} else {
		zs.callStateMap[callID] = state
	}
}

// Submit sends a call to Zammad, if we are aware of it and can correctly map its state to some known Zammad state
func (zs *Session) Submit(call *innovaphone.CallInSession) error {
	zs.Lock()
	defer zs.Unlock()

	src, dst := call.GetSourceAndDestination()
	dir := call.GetDirection()
	newState := call.GetState()

	curState := zs.get(call.Call)
	var setState State // used to only update our state if the call was submitted

	// only skip new calls, because a call-forwarding may send the call outside of the filtered group, meaning we have to
	// handle calls we would otherwise filter out based on involved numbers
	if curState == StateNew && !call.ShouldHandle() {
		logrus.WithField("call", call).Debugf("skipping call for number not in group '%s'", config.Global.PBX.FilterOnGroup)
		return nil
	}

	content := url.Values{
		"callId":    []string{strconv.Itoa(call.Call)},
		"from":      []string{normalizeNumber(src.E164)},
		"to":        []string{normalizeNumber(dst.E164)},
		"direction": []string{dir.String()},
	}

	var user string
	switch dir {
	case innovaphone.DirectionInbound:
		user = dst.Cn
	case innovaphone.DirectionOutbound:
		user = src.Cn
	}

	transition := stateTransition{curState: curState, newState: newState}
	switch transition {
	case
		stateTransition{StateNew, innovaphone.StateCallProc},
		stateTransition{StateNew, innovaphone.StateAlert}:
		// we do not get StateAlert on outgoing calls, so we have to assume StateCallProc is enough
		content.Set("event", "newCall")
		content.Set("user", user)
		setState = StateRinging
	case stateTransition{StateRinging, innovaphone.StateConnect}:
		content.Set("event", "answer")
		if dir == innovaphone.DirectionInbound {
			content.Set("user", user)
			content.Set("answeringNumber", normalizeNumber(dst.E164))
		}
		setState = StateConnected
	case stateTransition{StateRinging, innovaphone.StateDisconnectSent}:
		content.Set("event", "hangup")
		content.Set("cause", "cancel")
		setState = StateDisconnected
	case stateTransition{StateRinging, innovaphone.StateDisconnectReceived}:
		content.Set("event", "hangup")
		content.Set("cause", "noAnswer")
		setState = StateDisconnected
	case
		stateTransition{StateConnected, innovaphone.StateDisconnectSent},
		stateTransition{StateConnected, innovaphone.StateDisconnectReceived}:
		content.Set("event", "hangup")
		content.Set("cause", "normalClearing")
		setState = StateDisconnected
	default:
		logrus.WithField("call", call).Debugf("ignoring unknown state transition %s → %s", curState, newState)
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
	zs.update(call.Call, setState)
	return nil
}

func normalizeNumber(n string) string {
	if config.Global.Zammad.TrimFirstZero {
		return strings.TrimPrefix(n, "0")
	}
	return n
}

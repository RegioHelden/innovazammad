package innovaphone

import (
	"fmt"
	"strings"
)

// State represents the call state information from the innovaphone PBX' perspective.
type State int

// States as defined in the documentation (http://wiki.innovaphone.com/index.php?title=Reference10:SOAP_API)
const (
	_ State = iota // start at 1
	StateSetup
	StateSetupAck
	StateCallProc
	StateAlert
	StateConnect
	StateDisconnectSent
	StateDisconnectReceived
	StateParked
)

func (s State) String() string {
	switch s {
	case StateSetup:
		return "StateSetup"
	case StateSetupAck:
		return "StateSetupAck"
	case StateCallProc:
		return "StateCallProc"
	case StateAlert:
		return "StateAlert"
	case StateConnect:
		return "StateConnect"
	case StateDisconnectSent:
		return "StateDisconnectSent"
	case StateDisconnectReceived:
		return "StateDisconnectReceived"
	case StateParked:
		return "StateParked"
	default:
		panic("what!?")
	}
}

// Flip returns the state with directions switched, for those states which are direction-dependent.
func (s State) Flip() State {
	switch s {
	case StateDisconnectReceived:
		return StateDisconnectSent
	case StateDisconnectSent:
		return StateDisconnectReceived
	}
	return s
}

// Direction represents the call direction from the innovaphone PBX' perspective.
type Direction int

// Direction constants as defined in the documentation
const (
	DirectionInbound  = 0
	DirectionOutbound = 0x80
)

func (d Direction) String() string {
	switch d {
	case DirectionInbound:
		return "in"
	case DirectionOutbound:
		return "out"
	default:
		panic("what!?")
	}
}

// Flip returns the opposite diretion
func (d Direction) Flip() Direction {
	return (d ^ DirectionOutbound)
}

// Hold represents the call hold status
type Hold int

// Hold constants as defined in the documentation
const (
	HoldNone Hold = 0
	HoldThis Hold = 0x100
	HoldPeer Hold = 0x200
	HoldBoth Hold = 0x300
)

func (h Hold) String() string {
	switch h {
	case HoldNone:
		return "HoldNone"
	case HoldThis:
		return "HoldThis"
	case HoldPeer:
		return "HoldPeer"
	case HoldBoth:
		return "HoldBoth"
	default:
		panic("what!?")
	}
}

// GetSourceAndDestination returns the call's source and destination
func (call *CallInfo) GetSourceAndDestination() (src, dst *No) {
	for _, no := range call.No.Items {
		switch no.Type {
		case "this":
			dst = no
		case "peer":
			src = no
		}
	}

	if call.GetDirection() == DirectionInbound {
		src, dst = dst, src
	}

	return src, dst
}

// GetState returns the call's state
func (call *CallInfo) GetState() State {
	state := State(call.State & 0xF)
	if call.GetDirection() == DirectionInbound {
		state = state.Flip()
	}
	return state
}

// GetDirection returns the call's direction (wither incoming or outgoing)
func (call *CallInfo) GetDirection() Direction {
	return Direction(call.State & 0x80)
}

// GetHold returns the call's hold status (whether it was placed on hold by the local or remote part)
func (call *CallInfo) GetHold() Hold {
	return Hold(call.State & 0x300)
}

func (call *CallInfo) String() string {
	return fmt.Sprintf("CALL[%d][%s]", call.Call, call.No)
}

func (no *No) String() string {
	if no == nil {
		return "anonymous"
	}
	return no.E164
}

func (nos *NoArray) String() string {
	info := make([]string, len(nos.Items))
	for i, no := range nos.Items {
		if no.Cn != "" {
			info[i] = fmt.Sprintf("%s:%s (%s)", no.Type, no.E164, no.Cn)
		} else {
			info[i] = fmt.Sprintf("%s:%s", no.Type, no.E164)
		}
	}
	return strings.Join(info, ", ")
}

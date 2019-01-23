package innovaphone

import (
	"fmt"
	"strings"
)

type State int

const (
	StateSetup              State = 1 + iota
	StateSetupAck           State = 1 + iota
	StateCallProc           State = 1 + iota
	StateAlert              State = 1 + iota
	StateConnect            State = 1 + iota
	StateDisconnectSent     State = 1 + iota
	StateDisconnectReceived State = 1 + iota
	StateParked             State = 1 + iota
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

func (s State) Flip() State {
	switch s {
	case StateDisconnectReceived:
		return StateDisconnectSent
	case StateDisconnectSent:
		return StateDisconnectReceived
	}
	return s
}

type Direction int

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

func (d Direction) Flip() Direction {
	return (d ^ DirectionOutbound)
}

type Hold int

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

func (call *CallInfo) GetState() State {
	state := State(call.State & 0xF)
	if call.GetDirection() == DirectionInbound {
		state = state.Flip()
	}
	return state
}

func (call *CallInfo) GetDirection() Direction {
	return Direction(call.State & 0x80)
}

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

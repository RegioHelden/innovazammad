package innovaphone

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/regiohelden/innovazammad/config"
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
func (call *CallInfo) GetSourceAndDestination() (src, dst *No, user string, err error) {
	for _, no := range call.No {
		switch no.Type {
		case "this":
			dst = no
		case "peer":
			src = no
		}
	}

	if call.GetDirection() == DirectionInbound {
		src, dst = dst, src
		user = src.Cn
	} else {
		user = dst.Cn
	}

	if src == nil || dst == nil {
		return nil, nil, "", fmt.Errorf("call without src (%v) or dst (%v)", src, dst)
	}

	return src, dst, user, nil
}

// GetState returns the call's state
func (call *CallInfo) GetState() State {
	state := State(call.State & 0xF)
	if call.GetDirection() == DirectionInbound {
		state = state.Flip()
	}
	return state
}

// GetDirection returns the call's direction (either incoming or outgoing)
func (call *CallInfo) GetDirection() Direction {
	return Direction(call.State & 0x80)
}

// GetHold returns the call's hold status (whether it was placed on hold by the local or remote part)
func (call *CallInfo) GetHold() Hold {
	return Hold(call.State & 0x300)
}

func (call *CallInfo) String() string {
	return fmt.Sprintf("[%d][%s]", call.Call, call.No)
}

func (no *No) String() string {
	if no == nil {
		return "anonymous"
	}
	if no.Cn != "" {
		return fmt.Sprintf("%s:%s (%s)", no.Type, no.E164, no.Cn)
	}
	return fmt.Sprintf("%s:%s", no.Type, no.E164)
}

// Normalize returns an E123 formatted version of the number as a string.
// Local numbers are optionally prefixed with Config.Zammad.NumberPrefix, if they're not already in E123 format.
func (no *No) Normalize() (n string) {
	n = no.E164
	if config.Global.Zammad.TrunkPrefix != "" {
		n = strings.TrimPrefix(n, config.Global.Zammad.TrunkPrefix)
	}
	if strings.HasPrefix(n, "0") {
		n = fmt.Sprintf("%d%s", config.Global.Zammad.CountryCode, strings.TrimPrefix(n, "0"))
	} else if no.Type == "this" && !strings.HasPrefix(n, strconv.Itoa(config.Global.Zammad.CountryCode)) {
		n = fmt.Sprintf("%s%s", config.Global.Zammad.NumberPrefix, n)
	}
	return n
}

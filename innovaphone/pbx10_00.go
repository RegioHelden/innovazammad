package innovaphone

import (
	"reflect"

	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "http://innovaphone.com/pbx"

// NewPbxPortType creates an initializes a PbxPortType.
func NewPbxPortType(cli *soap.Client) PbxPortType {
	return &pbxPortType{cli}
}

// PbxPortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type PbxPortType interface {
	// Admin was auto-generated from WSDL.
	Admin(xml string) (string, error)

	// Calls was auto-generated from WSDL.
	Calls(session int, user string) (*CallInfoArray, error)

	// Devices was auto-generated from WSDL.
	Devices(session int, user string) (*DeviceArray, error)

	// Echo was auto-generated from WSDL.
	Echo(session int, key int) (int, error)

	// End was auto-generated from WSDL.
	End(session int) error

	// FindUser was auto-generated from WSDL.
	FindUser(v501 string, v700 string, v800 string, vx1000 string, cn string, h323 string, e164 string, count int, next int) (*UserInfoArray, error)

	// Initialize was auto-generated from WSDL.
	Initialize(user string, appl string, v bool, v501 bool, v700 bool, v800 bool, vx1000 bool) (int, int, error)

	// License was auto-generated from WSDL.
	License(session int, name string) (string, error)

	// LocationUrl was auto-generated from WSDL.
	LocationUrl(v501 string, v700 string, v800 string, vx1000 string, loc string) (string, error)

	// Poll was auto-generated from WSDL.
	Poll(session int) (*AnyInfo, error)

	// UserCall was auto-generated from WSDL.
	UserCall(user int, cn string, e164 string, h323 string, reg int, info *InfoArray, rc int, srce164 string) (int, error)

	// UserClear was auto-generated from WSDL.
	UserClear(call int, cause int, info *InfoArray) error

	// UserConnect was auto-generated from WSDL.
	UserConnect(call int) error

	// UserCtComplete was auto-generated from WSDL.
	UserCtComplete(call int, e164 string, h323 string) error

	// UserDTMF was auto-generated from WSDL.
	UserDTMF(call int, recv bool, dtmf string) error

	// UserEnd was auto-generated from WSDL.
	UserEnd(user int) error

	// UserFindDestination was auto-generated from WSDL.
	UserFindDestination(user int, e164 string, h323 string) (bool, *UserInfo, error)

	// UserHold was auto-generated from WSDL.
	UserHold(call int, remote bool) error

	// UserInfo was auto-generated from WSDL.
	UserInfo(call int, recv bool, cdpn string, key string, dsp string) error

	// UserInitialize was auto-generated from WSDL.
	UserInitialize(session int, user string, xfer bool, disc bool, hw string) (int, error)

	// UserLocalNum was auto-generated from WSDL.
	UserLocalNum(user int, num string) (string, error)

	// UserMediaTransfer was auto-generated from WSDL.
	UserMediaTransfer(a int, b int, user bool, peer bool) error

	// UserMessage was auto-generated from WSDL.
	UserMessage(user int, e164 string, h323 string, msg string, src_e164 string, src_h323 string) (int, error)

	// UserPark was auto-generated from WSDL.
	UserPark(call int, cn string, position int) (int, error)

	// UserPickup was auto-generated from WSDL.
	UserPickup(user int, cn string, call int, group string, reg int, info *InfoArray) (int, error)

	// UserRc was auto-generated from WSDL.
	UserRc(call int, rc int) error

	// UserRedirect was auto-generated from WSDL.
	UserRedirect(call int, cn string, e164 string, h323 string, info *InfoArray, rc int) (bool, error)

	// UserReroute was auto-generated from WSDL.
	UserReroute(call int, cn string, e164 string, h323 string, info *InfoArray) (bool, error)

	// UserRetrieve was auto-generated from WSDL.
	UserRetrieve(call int) error

	// UserTransfer was auto-generated from WSDL.
	UserTransfer(a int, b int) error

	// UserUUI was auto-generated from WSDL.
	UserUUI(call int, recv bool, uui string) error

	// Queries various version information from the PBX
	Version() (int, string, string, string, string, error)
}

// PresenceActivity was auto-generated from WSDL.
type PresenceActivity string

// Validate validates PresenceActivity.
func (v PresenceActivity) Validate() bool {
	for _, vv := range []string{
		"appointment",
		"away",
		"breakfast",
		"busy",
		"dinner",
		"holiday",
		"in-transit",
		"looking-for-work",
		"lunch",
		"meal",
		"meeting",
		"on-the-phone",
		"other",
		"performance",
		"permanent-absence",
		"playing",
		"presentation",
		"shopping",
		"sleeping",
		"spectator",
		"steering",
		"travel",
		"tv",
		"unknown",
		"vacation",
		"working",
		"worship",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// PresenceStatus was auto-generated from WSDL.
type PresenceStatus string

// Validate validates PresenceStatus.
func (v PresenceStatus) Validate() bool {
	for _, vv := range []string{
		"open",
		"closed",
	} {
		if reflect.DeepEqual(v, vv) {
			return true
		}
	}
	return false
}

// AnyInfo was auto-generated from WSDL.
type AnyInfo struct {
	User *UserInfoArray `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Call *CallInfoArray `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Reg  *RegInfoArray  `xml:"reg,omitempty" json:"reg,omitempty" yaml:"reg,omitempty"`
	Info *InfoArray     `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// CallInfo was auto-generated from WSDL.
type CallInfo struct {
	User   int        `xml:"user" json:"user" yaml:"user"`
	Call   int        `xml:"call" json:"call" yaml:"call"`
	Reg    int        `xml:"reg" json:"reg" yaml:"reg"`
	Active bool       `xml:"active" json:"active" yaml:"active"`
	State  int        `xml:"state" json:"state" yaml:"state"`
	No     *NoArray   `xml:"No" json:"No" yaml:"No"`
	Msg    string     `xml:"msg" json:"msg" yaml:"msg"`
	Info   *InfoArray `xml:"info" json:"info" yaml:"info"`
}

// CallInfoArray was auto-generated from WSDL.
type CallInfoArray struct {
	Items []*CallInfo `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
}

// Device was auto-generated from WSDL.
type Device struct {
	Hw   string `xml:"hw" json:"hw" yaml:"hw"`
	Text string `xml:"text" json:"text" yaml:"text"`
}

// DeviceArray was auto-generated from WSDL.
type DeviceArray struct {
	Items []*Device `xml:"device,omitempty" json:"device,omitempty" yaml:"device,omitempty"`
}

// Group was auto-generated from WSDL.
type Group struct {
	Group  string `xml:"group" json:"group" yaml:"group"`
	Active bool   `xml:"active" json:"active" yaml:"active"`
}

// GroupArray was auto-generated from WSDL.
type GroupArray struct {
	Items []*Group `xml:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
}

// Info was auto-generated from WSDL.
type Info struct {
	Type string `xml:"type" json:"type" yaml:"type"`
	Vali int    `xml:"vali" json:"vali" yaml:"vali"`
	Vals string `xml:"vals" json:"vals" yaml:"vals"`
}

// InfoArray was auto-generated from WSDL.
type InfoArray struct {
	Items []*Info `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// No was auto-generated from WSDL.
type No struct {
	Type string `xml:"type" json:"type" yaml:"type"`
	Cn   string `xml:"cn" json:"cn" yaml:"cn"`
	Dn   string `xml:"dn" json:"dn" yaml:"dn"`
	E164 string `xml:"e164" json:"e164" yaml:"e164"`
	H323 string `xml:"h323" json:"h323" yaml:"h323"`
}

// NoArray was auto-generated from WSDL.
type NoArray struct {
	Items []*No `xml:"No,omitempty" json:"No,omitempty" yaml:"No,omitempty"`
}

// Presence was auto-generated from WSDL.
type Presence struct {
	Status   PresenceStatus    `xml:"status" json:"status" yaml:"status"`
	Activity *PresenceActivity `xml:"activity,omitempty" json:"activity,omitempty" yaml:"activity,omitempty"`
	Note     *string           `xml:"note,omitempty" json:"note,omitempty" yaml:"note,omitempty"`
}

// RegInfo was auto-generated from WSDL.
type RegInfo struct {
	Active bool       `xml:"active" json:"active" yaml:"active"`
	User   int        `xml:"user" json:"user" yaml:"user"`
	Reg    int        `xml:"reg" json:"reg" yaml:"reg"`
	Hw     string     `xml:"hw" json:"hw" yaml:"hw"`
	Soap   string     `xml:"soap" json:"soap" yaml:"soap"`
	Info   *InfoArray `xml:"info" json:"info" yaml:"info"`
}

// RegInfoArray was auto-generated from WSDL.
type RegInfoArray struct {
	Items []*RegInfo `xml:"item,omitempty" json:"item,omitempty" yaml:"item,omitempty"`
}

// UserInfo was auto-generated from WSDL.
type UserInfo struct {
	Guid     string      `xml:"guid" json:"guid" yaml:"guid"`
	Active   bool        `xml:"active" json:"active" yaml:"active"`
	State    int         `xml:"state" json:"state" yaml:"state"`
	Channel  int         `xml:"channel" json:"channel" yaml:"channel"`
	Alert    int         `xml:"alert" json:"alert" yaml:"alert"`
	Cn       string      `xml:"cn" json:"cn" yaml:"cn"`
	Dn       string      `xml:"dn" json:"dn" yaml:"dn"`
	Type     string      `xml:"type" json:"type" yaml:"type"`
	E164     string      `xml:"e164" json:"e164" yaml:"e164"`
	H323     string      `xml:"h323" json:"h323" yaml:"h323"`
	Presence *Presence   `xml:"presence" json:"presence" yaml:"presence"`
	Loc      *string     `xml:"loc,omitempty" json:"loc,omitempty" yaml:"loc,omitempty"`
	Node     *string     `xml:"node,omitempty" json:"node,omitempty" yaml:"node,omitempty"`
	Nodenum  *string     `xml:"nodenum,omitempty" json:"nodenum,omitempty" yaml:"nodenum,omitempty"`
	Object   *string     `xml:"object,omitempty" json:"object,omitempty" yaml:"object,omitempty"`
	Cfg      *bool       `xml:"cfg,omitempty" json:"cfg,omitempty" yaml:"cfg,omitempty"`
	Groups   *GroupArray `xml:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Info     *InfoArray  `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// UserInfoArray was auto-generated from WSDL.
type UserInfoArray struct {
	Items []*UserInfo `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// Operation wrapper for Admin.
// OperationAdminRequest was auto-generated from WSDL.
type OperationAdminRequest struct {
	Xml *string `xml:"xml,omitempty" json:"xml,omitempty" yaml:"xml,omitempty"`
}

// Operation wrapper for Admin.
// OperationAdminResponse was auto-generated from WSDL.
type OperationAdminResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for Calls.
// OperationCallsRequest was auto-generated from WSDL.
type OperationCallsRequest struct {
	Session *int    `xml:"session,omitempty" json:"session,omitempty" yaml:"session,omitempty"`
	User    *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// Operation wrapper for Calls.
// OperationCallsResponse was auto-generated from WSDL.
type OperationCallsResponse struct {
	Return *CallInfoArray `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for Devices.
// OperationDevicesRequest was auto-generated from WSDL.
type OperationDevicesRequest struct {
	Session *int    `xml:"session,omitempty" json:"session,omitempty" yaml:"session,omitempty"`
	User    *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// Operation wrapper for Devices.
// OperationDevicesResponse was auto-generated from WSDL.
type OperationDevicesResponse struct {
	Return *DeviceArray `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for Echo.
// OperationEchoRequest was auto-generated from WSDL.
type OperationEchoRequest struct {
	Session *int `xml:"session,omitempty" json:"session,omitempty" yaml:"session,omitempty"`
	Key     *int `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Operation wrapper for Echo.
// OperationEchoResponse was auto-generated from WSDL.
type OperationEchoResponse struct {
	Return *int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for End.
// OperationEndRequest was auto-generated from WSDL.
type OperationEndRequest struct {
	Session *int `xml:"session,omitempty" json:"session,omitempty" yaml:"session,omitempty"`
}

// Operation wrapper for End.
// OperationEndResponse was auto-generated from WSDL.
type OperationEndResponse struct {
}

// Operation wrapper for FindUser.
// OperationFindUserRequest was auto-generated from WSDL.
type OperationFindUserRequest struct {
	V501   *string `xml:"v501,omitempty" json:"v501,omitempty" yaml:"v501,omitempty"`
	V700   *string `xml:"v700,omitempty" json:"v700,omitempty" yaml:"v700,omitempty"`
	V800   *string `xml:"v800,omitempty" json:"v800,omitempty" yaml:"v800,omitempty"`
	Vx1000 *string `xml:"vx1000,omitempty" json:"vx1000,omitempty" yaml:"vx1000,omitempty"`
	Cn     *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	H323   *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	E164   *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	Count  *int    `xml:"count,omitempty" json:"count,omitempty" yaml:"count,omitempty"`
	Next   *int    `xml:"next,omitempty" json:"next,omitempty" yaml:"next,omitempty"`
}

// Operation wrapper for FindUser.
// OperationFindUserResponse was auto-generated from WSDL.
type OperationFindUserResponse struct {
	Return *UserInfoArray `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for Initialize.
// OperationInitializeRequest was auto-generated from WSDL.
type OperationInitializeRequest struct {
	User   *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Appl   *string `xml:"appl,omitempty" json:"appl,omitempty" yaml:"appl,omitempty"`
	V      *bool   `xml:"v,omitempty" json:"v,omitempty" yaml:"v,omitempty"`
	V501   *bool   `xml:"v501,omitempty" json:"v501,omitempty" yaml:"v501,omitempty"`
	V700   *bool   `xml:"v700,omitempty" json:"v700,omitempty" yaml:"v700,omitempty"`
	V800   *bool   `xml:"v800,omitempty" json:"v800,omitempty" yaml:"v800,omitempty"`
	Vx1000 *bool   `xml:"vx1000,omitempty" json:"vx1000,omitempty" yaml:"vx1000,omitempty"`
}

// Operation wrapper for Initialize.
// OperationInitializeResponse was auto-generated from WSDL.
type OperationInitializeResponse struct {
	Return *int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
	Key    *int `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
}

// Operation wrapper for License.
// OperationLicenseRequest was auto-generated from WSDL.
type OperationLicenseRequest struct {
	Session *int    `xml:"session,omitempty" json:"session,omitempty" yaml:"session,omitempty"`
	Name    *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// Operation wrapper for License.
// OperationLicenseResponse was auto-generated from WSDL.
type OperationLicenseResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for LocationUrl.
// OperationLocationUrlRequest was auto-generated from WSDL.
type OperationLocationUrlRequest struct {
	V501   *string `xml:"v501,omitempty" json:"v501,omitempty" yaml:"v501,omitempty"`
	V700   *string `xml:"v700,omitempty" json:"v700,omitempty" yaml:"v700,omitempty"`
	V800   *string `xml:"v800,omitempty" json:"v800,omitempty" yaml:"v800,omitempty"`
	Vx1000 *string `xml:"vx1000,omitempty" json:"vx1000,omitempty" yaml:"vx1000,omitempty"`
	Loc    *string `xml:"loc,omitempty" json:"loc,omitempty" yaml:"loc,omitempty"`
}

// Operation wrapper for LocationUrl.
// OperationLocationUrlResponse was auto-generated from WSDL.
type OperationLocationUrlResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for Poll.
// OperationPollRequest was auto-generated from WSDL.
type OperationPollRequest struct {
	Session *int `xml:"session,omitempty" json:"session,omitempty" yaml:"session,omitempty"`
}

// Operation wrapper for Poll.
// OperationPollResponse was auto-generated from WSDL.
type OperationPollResponse struct {
	Return *AnyInfo `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserCall.
// OperationUserCallRequest was auto-generated from WSDL.
type OperationUserCallRequest struct {
	User    *int       `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Cn      *string    `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	E164    *string    `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323    *string    `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Reg     *int       `xml:"reg,omitempty" json:"reg,omitempty" yaml:"reg,omitempty"`
	Info    *InfoArray `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
	Rc      *int       `xml:"rc,omitempty" json:"rc,omitempty" yaml:"rc,omitempty"`
	Srce164 *string    `xml:"srce164,omitempty" json:"srce164,omitempty" yaml:"srce164,omitempty"`
}

// Operation wrapper for UserCall.
// OperationUserCallResponse was auto-generated from WSDL.
type OperationUserCallResponse struct {
	Return *int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserClear.
// OperationUserClearRequest was auto-generated from WSDL.
type OperationUserClearRequest struct {
	Call  *int       `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Cause *int       `xml:"cause,omitempty" json:"cause,omitempty" yaml:"cause,omitempty"`
	Info  *InfoArray `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for UserClear.
// OperationUserClearResponse was auto-generated from WSDL.
type OperationUserClearResponse struct {
}

// Operation wrapper for UserConnect.
// OperationUserConnectRequest was auto-generated from WSDL.
type OperationUserConnectRequest struct {
	Call *int `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
}

// Operation wrapper for UserConnect.
// OperationUserConnectResponse was auto-generated from WSDL.
type OperationUserConnectResponse struct {
}

// Operation wrapper for UserCtComplete.
// OperationUserCtCompleteRequest was auto-generated from WSDL.
type OperationUserCtCompleteRequest struct {
	Call *int    `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	E164 *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
}

// Operation wrapper for UserCtComplete.
// OperationUserCtCompleteResponse was auto-generated from WSDL.
type OperationUserCtCompleteResponse struct {
}

// Operation wrapper for UserDTMF.
// OperationUserDTMFRequest was auto-generated from WSDL.
type OperationUserDTMFRequest struct {
	Call *int    `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Recv *bool   `xml:"recv,omitempty" json:"recv,omitempty" yaml:"recv,omitempty"`
	Dtmf *string `xml:"dtmf,omitempty" json:"dtmf,omitempty" yaml:"dtmf,omitempty"`
}

// Operation wrapper for UserDTMF.
// OperationUserDTMFResponse was auto-generated from WSDL.
type OperationUserDTMFResponse struct {
}

// Operation wrapper for UserEnd.
// OperationUserEndRequest was auto-generated from WSDL.
type OperationUserEndRequest struct {
	User *int `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// Operation wrapper for UserEnd.
// OperationUserEndResponse was auto-generated from WSDL.
type OperationUserEndResponse struct {
}

// Operation wrapper for UserFindDestination.
// OperationUserFindDestinationRequest was auto-generated from
// WSDL.
type OperationUserFindDestinationRequest struct {
	User *int    `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	E164 *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
}

// Operation wrapper for UserFindDestination.
// OperationUserFindDestinationResponse was auto-generated from
// WSDL.
type OperationUserFindDestinationResponse struct {
	Return *bool     `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
	User   *UserInfo `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// Operation wrapper for UserHold.
// OperationUserHoldRequest was auto-generated from WSDL.
type OperationUserHoldRequest struct {
	Call   *int  `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Remote *bool `xml:"remote,omitempty" json:"remote,omitempty" yaml:"remote,omitempty"`
}

// Operation wrapper for UserHold.
// OperationUserHoldResponse was auto-generated from WSDL.
type OperationUserHoldResponse struct {
}

// Operation wrapper for UserInfo.
// OperationUserInfoRequest was auto-generated from WSDL.
type OperationUserInfoRequest struct {
	Call *int    `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Recv *bool   `xml:"recv,omitempty" json:"recv,omitempty" yaml:"recv,omitempty"`
	Cdpn *string `xml:"cdpn,omitempty" json:"cdpn,omitempty" yaml:"cdpn,omitempty"`
	Key  *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Dsp  *string `xml:"dsp,omitempty" json:"dsp,omitempty" yaml:"dsp,omitempty"`
}

// Operation wrapper for UserInfo.
// OperationUserInfoResponse was auto-generated from WSDL.
type OperationUserInfoResponse struct {
}

// Operation wrapper for UserInitialize.
// OperationUserInitializeRequest was auto-generated from WSDL.
type OperationUserInitializeRequest struct {
	Session *int    `xml:"session,omitempty" json:"session,omitempty" yaml:"session,omitempty"`
	User    *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Xfer    *bool   `xml:"xfer,omitempty" json:"xfer,omitempty" yaml:"xfer,omitempty"`
	Disc    *bool   `xml:"disc,omitempty" json:"disc,omitempty" yaml:"disc,omitempty"`
	Hw      *string `xml:"hw,omitempty" json:"hw,omitempty" yaml:"hw,omitempty"`
}

// Operation wrapper for UserInitialize.
// OperationUserInitializeResponse was auto-generated from WSDL.
type OperationUserInitializeResponse struct {
	Return *int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserLocalNum.
// OperationUserLocalNumRequest was auto-generated from WSDL.
type OperationUserLocalNumRequest struct {
	User *int    `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Num  *string `xml:"num,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
}

// Operation wrapper for UserLocalNum.
// OperationUserLocalNumResponse was auto-generated from WSDL.
type OperationUserLocalNumResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserMediaTransfer.
// OperationUserMediaTransferRequest was auto-generated from WSDL.
type OperationUserMediaTransferRequest struct {
	A    *int  `xml:"a,omitempty" json:"a,omitempty" yaml:"a,omitempty"`
	B    *int  `xml:"b,omitempty" json:"b,omitempty" yaml:"b,omitempty"`
	User *bool `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Peer *bool `xml:"peer,omitempty" json:"peer,omitempty" yaml:"peer,omitempty"`
}

// Operation wrapper for UserMediaTransfer.
// OperationUserMediaTransferResponse was auto-generated from WSDL.
type OperationUserMediaTransferResponse struct {
}

// Operation wrapper for UserMessage.
// OperationUserMessageRequest was auto-generated from WSDL.
type OperationUserMessageRequest struct {
	User     *int    `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	E164     *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323     *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Msg      *string `xml:"msg,omitempty" json:"msg,omitempty" yaml:"msg,omitempty"`
	Src_e164 *string `xml:"src_e164,omitempty" json:"src_e164,omitempty" yaml:"src_e164,omitempty"`
	Src_h323 *string `xml:"src_h323,omitempty" json:"src_h323,omitempty" yaml:"src_h323,omitempty"`
}

// Operation wrapper for UserMessage.
// OperationUserMessageResponse was auto-generated from WSDL.
type OperationUserMessageResponse struct {
	Return *int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserPark.
// OperationUserParkRequest was auto-generated from WSDL.
type OperationUserParkRequest struct {
	Call     *int    `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Cn       *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	Position *int    `xml:"position,omitempty" json:"position,omitempty" yaml:"position,omitempty"`
}

// Operation wrapper for UserPark.
// OperationUserParkResponse was auto-generated from WSDL.
type OperationUserParkResponse struct {
	Return *int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserPickup.
// OperationUserPickupRequest was auto-generated from WSDL.
type OperationUserPickupRequest struct {
	User  *int       `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Cn    *string    `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	Call  *int       `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Group *string    `xml:"group,omitempty" json:"group,omitempty" yaml:"group,omitempty"`
	Reg   *int       `xml:"reg,omitempty" json:"reg,omitempty" yaml:"reg,omitempty"`
	Info  *InfoArray `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for UserPickup.
// OperationUserPickupResponse was auto-generated from WSDL.
type OperationUserPickupResponse struct {
	Return *int `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserRc.
// OperationUserRcRequest was auto-generated from WSDL.
type OperationUserRcRequest struct {
	Call *int `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Rc   *int `xml:"rc,omitempty" json:"rc,omitempty" yaml:"rc,omitempty"`
}

// Operation wrapper for UserRc.
// OperationUserRcResponse was auto-generated from WSDL.
type OperationUserRcResponse struct {
}

// Operation wrapper for UserRedirect.
// OperationUserRedirectRequest was auto-generated from WSDL.
type OperationUserRedirectRequest struct {
	Call *int       `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Cn   *string    `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	E164 *string    `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string    `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Info *InfoArray `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
	Rc   *int       `xml:"rc,omitempty" json:"rc,omitempty" yaml:"rc,omitempty"`
}

// Operation wrapper for UserRedirect.
// OperationUserRedirectResponse was auto-generated from WSDL.
type OperationUserRedirectResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserReroute.
// OperationUserRerouteRequest was auto-generated from WSDL.
type OperationUserRerouteRequest struct {
	Call *int       `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Cn   *string    `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	E164 *string    `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string    `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Info *InfoArray `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Operation wrapper for UserReroute.
// OperationUserRerouteResponse was auto-generated from WSDL.
type OperationUserRerouteResponse struct {
	Return *bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for UserRetrieve.
// OperationUserRetrieveRequest was auto-generated from WSDL.
type OperationUserRetrieveRequest struct {
	Call *int `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
}

// Operation wrapper for UserRetrieve.
// OperationUserRetrieveResponse was auto-generated from WSDL.
type OperationUserRetrieveResponse struct {
}

// Operation wrapper for UserTransfer.
// OperationUserTransferRequest was auto-generated from WSDL.
type OperationUserTransferRequest struct {
	A *int `xml:"a,omitempty" json:"a,omitempty" yaml:"a,omitempty"`
	B *int `xml:"b,omitempty" json:"b,omitempty" yaml:"b,omitempty"`
}

// Operation wrapper for UserTransfer.
// OperationUserTransferResponse was auto-generated from WSDL.
type OperationUserTransferResponse struct {
}

// Operation wrapper for UserUUI.
// OperationUserUUIRequest was auto-generated from WSDL.
type OperationUserUUIRequest struct {
	Call *int    `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Recv *bool   `xml:"recv,omitempty" json:"recv,omitempty" yaml:"recv,omitempty"`
	Uui  *string `xml:"uui,omitempty" json:"uui,omitempty" yaml:"uui,omitempty"`
}

// Operation wrapper for UserUUI.
// OperationUserUUIResponse was auto-generated from WSDL.
type OperationUserUUIResponse struct {
}

// Operation wrapper for Version.
// OperationVersionResponse was auto-generated from WSDL.
type OperationVersionResponse struct {
	WSDLVersion     *int    `xml:"WSDLVersion,omitempty" json:"WSDLVersion,omitempty" yaml:"WSDLVersion,omitempty"`
	GatekeeperID    *string `xml:"GatekeeperID,omitempty" json:"GatekeeperID,omitempty" yaml:"GatekeeperID,omitempty"`
	Location        *string `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	FirmwareVersion *string `xml:"FirmwareVersion,omitempty" json:"FirmwareVersion,omitempty" yaml:"FirmwareVersion,omitempty"`
	SerialNumber    *string `xml:"SerialNumber,omitempty" json:"SerialNumber,omitempty" yaml:"SerialNumber,omitempty"`
}

// pbxPortType implements the PbxPortType interface.
type pbxPortType struct {
	cli *soap.Client
}

// Admin was auto-generated from WSDL.
func (p *pbxPortType) Admin(xml string) (string, error) {
	α := struct {
		M OperationAdminRequest `xml:"inno:Admin"`
	}{
		OperationAdminRequest{
			&xml,
		},
	}

	γ := struct {
		M OperationAdminResponse `xml:"AdminResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#Admin", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Calls was auto-generated from WSDL.
func (p *pbxPortType) Calls(session int, user string) (*CallInfoArray, error) {
	α := struct {
		M OperationCallsRequest `xml:"inno:Calls"`
	}{
		OperationCallsRequest{
			&session,
			&user,
		},
	}

	γ := struct {
		M OperationCallsResponse `xml:"CallsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#Calls", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// Devices was auto-generated from WSDL.
func (p *pbxPortType) Devices(session int, user string) (*DeviceArray, error) {
	α := struct {
		M OperationDevicesRequest `xml:"inno:Devices"`
	}{
		OperationDevicesRequest{
			&session,
			&user,
		},
	}

	γ := struct {
		M OperationDevicesResponse `xml:"DevicesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#Devices", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// Echo was auto-generated from WSDL.
func (p *pbxPortType) Echo(session int, key int) (int, error) {
	α := struct {
		M OperationEchoRequest `xml:"inno:Echo"`
	}{
		OperationEchoRequest{
			&session,
			&key,
		},
	}

	γ := struct {
		M OperationEchoResponse `xml:"EchoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#Echo", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Return, nil
}

// End was auto-generated from WSDL.
func (p *pbxPortType) End(session int) error {
	α := struct {
		M OperationEndRequest `xml:"inno:End"`
	}{
		OperationEndRequest{
			&session,
		},
	}

	γ := struct {
		M OperationEndResponse `xml:"EndResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#End", α, &γ); err != nil {
		return err
	}
	return nil
}

// FindUser was auto-generated from WSDL.
func (p *pbxPortType) FindUser(v501 string, v700 string, v800 string, vx1000 string, cn string, h323 string, e164 string, count int, next int) (*UserInfoArray, error) {
	α := struct {
		M OperationFindUserRequest `xml:"inno:FindUser"`
	}{
		OperationFindUserRequest{
			&v501,
			&v700,
			&v800,
			&vx1000,
			&cn,
			&h323,
			&e164,
			&count,
			&next,
		},
	}

	γ := struct {
		M OperationFindUserResponse `xml:"FindUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#FindUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// Initialize was auto-generated from WSDL.
func (p *pbxPortType) Initialize(user string, appl string, v bool, v501 bool, v700 bool, v800 bool, vx1000 bool) (int, int, error) {
	α := struct {
		M OperationInitializeRequest `xml:"inno:Initialize"`
	}{
		OperationInitializeRequest{
			&user,
			&appl,
			&v,
			&v501,
			&v700,
			&v800,
			&vx1000,
		},
	}

	γ := struct {
		M OperationInitializeResponse `xml:"InitializeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#Initialize", α, &γ); err != nil {
		return 0, 0, err
	}
	return *γ.M.Return, *γ.M.Key, nil
}

// License was auto-generated from WSDL.
func (p *pbxPortType) License(session int, name string) (string, error) {
	α := struct {
		M OperationLicenseRequest `xml:"inno:License"`
	}{
		OperationLicenseRequest{
			&session,
			&name,
		},
	}

	γ := struct {
		M OperationLicenseResponse `xml:"LicenseResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#License", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// LocationUrl was auto-generated from WSDL.
func (p *pbxPortType) LocationUrl(v501 string, v700 string, v800 string, vx1000 string, loc string) (string, error) {
	α := struct {
		M OperationLocationUrlRequest `xml:"inno:LocationUrl"`
	}{
		OperationLocationUrlRequest{
			&v501,
			&v700,
			&v800,
			&vx1000,
			&loc,
		},
	}

	γ := struct {
		M OperationLocationUrlResponse `xml:"LocationUrlResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#LocationUrl", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Poll was auto-generated from WSDL.
func (p *pbxPortType) Poll(session int) (*AnyInfo, error) {
	α := struct {
		M OperationPollRequest `xml:"inno:Poll"`
	}{
		OperationPollRequest{
			&session,
		},
	}

	γ := struct {
		M OperationPollResponse `xml:"PollResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#Poll", α, &γ); err != nil {
		return nil, err
	}
	return γ.M.Return, nil
}

// UserCall was auto-generated from WSDL.
func (p *pbxPortType) UserCall(user int, cn string, e164 string, h323 string, reg int, info *InfoArray, rc int, srce164 string) (int, error) {
	α := struct {
		M OperationUserCallRequest `xml:"inno:UserCall"`
	}{
		OperationUserCallRequest{
			&user,
			&cn,
			&e164,
			&h323,
			&reg,
			info,
			&rc,
			&srce164,
		},
	}

	γ := struct {
		M OperationUserCallResponse `xml:"UserCallResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserCall", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Return, nil
}

// UserClear was auto-generated from WSDL.
func (p *pbxPortType) UserClear(call int, cause int, info *InfoArray) error {
	α := struct {
		M OperationUserClearRequest `xml:"inno:UserClear"`
	}{
		OperationUserClearRequest{
			&call,
			&cause,
			info,
		},
	}

	γ := struct {
		M OperationUserClearResponse `xml:"UserClearResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserClear", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserConnect was auto-generated from WSDL.
func (p *pbxPortType) UserConnect(call int) error {
	α := struct {
		M OperationUserConnectRequest `xml:"inno:UserConnect"`
	}{
		OperationUserConnectRequest{
			&call,
		},
	}

	γ := struct {
		M OperationUserConnectResponse `xml:"UserConnectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserConnect", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserCtComplete was auto-generated from WSDL.
func (p *pbxPortType) UserCtComplete(call int, e164 string, h323 string) error {
	α := struct {
		M OperationUserCtCompleteRequest `xml:"inno:UserCtComplete"`
	}{
		OperationUserCtCompleteRequest{
			&call,
			&e164,
			&h323,
		},
	}

	γ := struct {
		M OperationUserCtCompleteResponse `xml:"UserCtCompleteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserCtComplete", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserDTMF was auto-generated from WSDL.
func (p *pbxPortType) UserDTMF(call int, recv bool, dtmf string) error {
	α := struct {
		M OperationUserDTMFRequest `xml:"inno:UserDTMF"`
	}{
		OperationUserDTMFRequest{
			&call,
			&recv,
			&dtmf,
		},
	}

	γ := struct {
		M OperationUserDTMFResponse `xml:"UserDTMFResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserDTMF", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserEnd was auto-generated from WSDL.
func (p *pbxPortType) UserEnd(user int) error {
	α := struct {
		M OperationUserEndRequest `xml:"inno:UserEnd"`
	}{
		OperationUserEndRequest{
			&user,
		},
	}

	γ := struct {
		M OperationUserEndResponse `xml:"UserEndResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("capeconnect:pbx:pbxPortType#UserEnd", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserFindDestination was auto-generated from WSDL.
func (p *pbxPortType) UserFindDestination(user int, e164 string, h323 string) (bool, *UserInfo, error) {
	α := struct {
		M OperationUserFindDestinationRequest `xml:"inno:UserFindDestination"`
	}{
		OperationUserFindDestinationRequest{
			&user,
			&e164,
			&h323,
		},
	}

	γ := struct {
		M OperationUserFindDestinationResponse `xml:"UserFindDestinationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserFindDestination", α, &γ); err != nil {
		return false, nil, err
	}
	return *γ.M.Return, γ.M.User, nil
}

// UserHold was auto-generated from WSDL.
func (p *pbxPortType) UserHold(call int, remote bool) error {
	α := struct {
		M OperationUserHoldRequest `xml:"inno:UserHold"`
	}{
		OperationUserHoldRequest{
			&call,
			&remote,
		},
	}

	γ := struct {
		M OperationUserHoldResponse `xml:"UserHoldResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserHold", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserInfo was auto-generated from WSDL.
func (p *pbxPortType) UserInfo(call int, recv bool, cdpn string, key string, dsp string) error {
	α := struct {
		M OperationUserInfoRequest `xml:"inno:UserInfo"`
	}{
		OperationUserInfoRequest{
			&call,
			&recv,
			&cdpn,
			&key,
			&dsp,
		},
	}

	γ := struct {
		M OperationUserInfoResponse `xml:"UserInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserInfo", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserInitialize was auto-generated from WSDL.
func (p *pbxPortType) UserInitialize(session int, user string, xfer bool, disc bool, hw string) (int, error) {
	α := struct {
		M OperationUserInitializeRequest `xml:"inno:UserInitialize"`
	}{
		OperationUserInitializeRequest{
			&session,
			&user,
			&xfer,
			&disc,
			&hw,
		},
	}

	γ := struct {
		M OperationUserInitializeResponse `xml:"UserInitializeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserInitialize", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Return, nil
}

// UserLocalNum was auto-generated from WSDL.
func (p *pbxPortType) UserLocalNum(user int, num string) (string, error) {
	α := struct {
		M OperationUserLocalNumRequest `xml:"inno:UserLocalNum"`
	}{
		OperationUserLocalNumRequest{
			&user,
			&num,
		},
	}

	γ := struct {
		M OperationUserLocalNumResponse `xml:"UserLocalNumResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserLocalNum", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// UserMediaTransfer was auto-generated from WSDL.
func (p *pbxPortType) UserMediaTransfer(a int, b int, user bool, peer bool) error {
	α := struct {
		M OperationUserMediaTransferRequest `xml:"inno:UserMediaTransfer"`
	}{
		OperationUserMediaTransferRequest{
			&a,
			&b,
			&user,
			&peer,
		},
	}

	γ := struct {
		M OperationUserMediaTransferResponse `xml:"UserMediaTransferResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserMediaTransfer", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserMessage was auto-generated from WSDL.
func (p *pbxPortType) UserMessage(user int, e164 string, h323 string, msg string, src_e164 string, src_h323 string) (int, error) {
	α := struct {
		M OperationUserMessageRequest `xml:"inno:UserMessage"`
	}{
		OperationUserMessageRequest{
			&user,
			&e164,
			&h323,
			&msg,
			&src_e164,
			&src_h323,
		},
	}

	γ := struct {
		M OperationUserMessageResponse `xml:"UserMessageResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserMessage", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Return, nil
}

// UserPark was auto-generated from WSDL.
func (p *pbxPortType) UserPark(call int, cn string, position int) (int, error) {
	α := struct {
		M OperationUserParkRequest `xml:"inno:UserPark"`
	}{
		OperationUserParkRequest{
			&call,
			&cn,
			&position,
		},
	}

	γ := struct {
		M OperationUserParkResponse `xml:"UserParkResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserPark", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Return, nil
}

// UserPickup was auto-generated from WSDL.
func (p *pbxPortType) UserPickup(user int, cn string, call int, group string, reg int, info *InfoArray) (int, error) {
	α := struct {
		M OperationUserPickupRequest `xml:"inno:UserPickup"`
	}{
		OperationUserPickupRequest{
			&user,
			&cn,
			&call,
			&group,
			&reg,
			info,
		},
	}

	γ := struct {
		M OperationUserPickupResponse `xml:"UserPickupResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserPickup", α, &γ); err != nil {
		return 0, err
	}
	return *γ.M.Return, nil
}

// UserRc was auto-generated from WSDL.
func (p *pbxPortType) UserRc(call int, rc int) error {
	α := struct {
		M OperationUserRcRequest `xml:"inno:UserRc"`
	}{
		OperationUserRcRequest{
			&call,
			&rc,
		},
	}

	γ := struct {
		M OperationUserRcResponse `xml:"UserRcResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserRc", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserRedirect was auto-generated from WSDL.
func (p *pbxPortType) UserRedirect(call int, cn string, e164 string, h323 string, info *InfoArray, rc int) (bool, error) {
	α := struct {
		M OperationUserRedirectRequest `xml:"inno:UserRedirect"`
	}{
		OperationUserRedirectRequest{
			&call,
			&cn,
			&e164,
			&h323,
			info,
			&rc,
		},
	}

	γ := struct {
		M OperationUserRedirectResponse `xml:"UserRedirectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserRedirect", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Return, nil
}

// UserReroute was auto-generated from WSDL.
func (p *pbxPortType) UserReroute(call int, cn string, e164 string, h323 string, info *InfoArray) (bool, error) {
	α := struct {
		M OperationUserRerouteRequest `xml:"inno:UserReroute"`
	}{
		OperationUserRerouteRequest{
			&call,
			&cn,
			&e164,
			&h323,
			info,
		},
	}

	γ := struct {
		M OperationUserRerouteResponse `xml:"UserRerouteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserReroute", α, &γ); err != nil {
		return false, err
	}
	return *γ.M.Return, nil
}

// UserRetrieve was auto-generated from WSDL.
func (p *pbxPortType) UserRetrieve(call int) error {
	α := struct {
		M OperationUserRetrieveRequest `xml:"inno:UserRetrieve"`
	}{
		OperationUserRetrieveRequest{
			&call,
		},
	}

	γ := struct {
		M OperationUserRetrieveResponse `xml:"UserRetrieveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserRetrieve", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserTransfer was auto-generated from WSDL.
func (p *pbxPortType) UserTransfer(a int, b int) error {
	α := struct {
		M OperationUserTransferRequest `xml:"inno:UserTransfer"`
	}{
		OperationUserTransferRequest{
			&a,
			&b,
		},
	}

	γ := struct {
		M OperationUserTransferResponse `xml:"UserTransferResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserTransfer", α, &γ); err != nil {
		return err
	}
	return nil
}

// UserUUI was auto-generated from WSDL.
func (p *pbxPortType) UserUUI(call int, recv bool, uui string) error {
	α := struct {
		M OperationUserUUIRequest `xml:"inno:UserUUI"`
	}{
		OperationUserUUIRequest{
			&call,
			&recv,
			&uui,
		},
	}

	γ := struct {
		M OperationUserUUIResponse `xml:"UserUUIResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#UserUUI", α, &γ); err != nil {
		return err
	}
	return nil
}

// Queries various version information from the PBX
func (p *pbxPortType) Version() (int, string, string, string, string, error) {
	α := struct {
		M struct{} `xml:"inno:Version"`
	}{
		struct{}{},
	}

	γ := struct {
		M OperationVersionResponse `xml:"VersionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx#Version", α, &γ); err != nil {
		return 0, "", "", "", "", err
	}
	return *γ.M.WSDLVersion, *γ.M.GatekeeperID, *γ.M.Location, *γ.M.FirmwareVersion, *γ.M.SerialNumber, nil
}

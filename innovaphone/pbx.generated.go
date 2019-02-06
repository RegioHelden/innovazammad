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
	Admin(Admin *Admin) (*AdminResponse, error)

	// Calls was auto-generated from WSDL.
	Calls(Calls *Calls) (*CallsResponse, error)

	// Devices was auto-generated from WSDL.
	Devices(Devices *Devices) (*DevicesResponse, error)

	// Echo was auto-generated from WSDL.
	Echo(Echo *Echo) (*EchoResponse, error)

	// End was auto-generated from WSDL.
	End(End *End) (*EndResponse, error)

	// FindUser was auto-generated from WSDL.
	FindUser(FindUser *FindUser) (*FindUserResponse, error)

	// Initialize was auto-generated from WSDL.
	Initialize(Initialize *Initialize) (*InitializeResponse, error)

	// License was auto-generated from WSDL.
	License(License *License) (*LicenseResponse, error)

	// LocationUrl was auto-generated from WSDL.
	LocationUrl(LocationUrl *LocationUrl) (*LocationUrlResponse, error)

	// Poll was auto-generated from WSDL.
	Poll(Poll *Poll) (*PollResponse, error)

	// SetPresence was auto-generated from WSDL.
	SetPresence(SetPresence *SetPresence) (*SetPresenceResponse, error)

	// UserCall was auto-generated from WSDL.
	UserCall(UserCall *UserCall) (*UserCallResponse, error)

	// UserClear was auto-generated from WSDL.
	UserClear(UserClear *UserClear) (*UserClearResponse, error)

	// UserConnect was auto-generated from WSDL.
	UserConnect(UserConnect *UserConnect) (*UserConnectResponse, error)

	// UserCtComplete was auto-generated from WSDL.
	UserCtComplete(UserCtComplete *UserCtComplete) (*UserCtCompleteResponse, error)

	// UserDTMF was auto-generated from WSDL.
	UserDTMF(UserDTMF *UserDTMF) (*UserDTMFResponse, error)

	// UserEnd was auto-generated from WSDL.
	UserEnd(UserEnd *UserEnd) (*UserEndResponse, error)

	// UserFindDestination was auto-generated from WSDL.
	UserFindDestination(UserFindDestination *UserFindDestination) (*UserFindDestinationResponse, error)

	// UserHold was auto-generated from WSDL.
	UserHold(UserHold *UserHold) (*UserHoldResponse, error)

	// UserInfo was auto-generated from WSDL.
	UserInfo(UserInfoTx *UserInfoTx) (*UserInfoTxResponse, error)

	// UserInitialize was auto-generated from WSDL.
	UserInitialize(UserInitialize *UserInitialize) (*UserInitializeResponse, error)

	// UserLicense was auto-generated from WSDL.
	UserLicense(UserLicense *UserLicense) (*UserLicenseResponse, error)

	// UserLocalNum was auto-generated from WSDL.
	UserLocalNum(UserLocalNum *UserLocalNum) (*UserLocalNumResponse, error)

	// UserMediaTransfer was auto-generated from WSDL.
	UserMediaTransfer(UserMediaTransfer *UserMediaTransfer) (*UserMediaTransferResponse, error)

	// UserMessage was auto-generated from WSDL.
	UserMessage(UserMessage *UserMessage) (*UserMessageResponse, error)

	// UserPark was auto-generated from WSDL.
	UserPark(UserPark *UserPark) (*UserParkResponse, error)

	// UserPickup was auto-generated from WSDL.
	UserPickup(UserPickup *UserPickup) (*UserPickupResponse, error)

	// UserRc was auto-generated from WSDL.
	UserRc(UserRc *UserRc) (*UserRcResponse, error)

	// UserRedirect was auto-generated from WSDL.
	UserRedirect(UserRedirect *UserRedirect) (*UserRedirectResponse, error)

	// UserReroute was auto-generated from WSDL.
	UserReroute(UserReroute *UserReroute) (*UserRerouteResponse, error)

	// UserRetrieve was auto-generated from WSDL.
	UserRetrieve(UserRetrieve *UserRetrieve) (*UserRetrieveResponse, error)

	// UserTransfer was auto-generated from WSDL.
	UserTransfer(UserTransfer *UserTransfer) (*UserTransferResponse, error)

	// UserUUI was auto-generated from WSDL.
	UserUUI(UserUUI *UserUUI) (*UserUUIResponse, error)

	// Version was auto-generated from WSDL.
	Version(Version *Version) (*VersionResponse, error)
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
		"dnd",
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

// Admin was auto-generated from WSDL.
type Admin struct {
	Xml *string `xml:"xml,omitempty" json:"xml,omitempty" yaml:"xml,omitempty"`
}

// AdminResponse was auto-generated from WSDL.
type AdminResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// AnyInfo was auto-generated from WSDL.
type AnyInfo struct {
	User []*UserInfo `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Call []*CallInfo `xml:"call,omitempty" json:"call,omitempty" yaml:"call,omitempty"`
	Reg  []*RegInfo  `xml:"reg,omitempty" json:"reg,omitempty" yaml:"reg,omitempty"`
	Info []*Info     `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// CallInfo was auto-generated from WSDL.
type CallInfo struct {
	User   int     `xml:"user" json:"user" yaml:"user"`
	Call   int     `xml:"call" json:"call" yaml:"call"`
	Reg    int     `xml:"reg" json:"reg" yaml:"reg"`
	Active bool    `xml:"active" json:"active" yaml:"active"`
	State  int     `xml:"state" json:"state" yaml:"state"`
	No     []*No   `xml:"No,omitempty" json:"No,omitempty" yaml:"No,omitempty"`
	Msg    string  `xml:"msg" json:"msg" yaml:"msg"`
	Info   []*Info `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// Calls was auto-generated from WSDL.
type Calls struct {
	Session int     `xml:"session" json:"session" yaml:"session"`
	User    *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// CallsResponse was auto-generated from WSDL.
type CallsResponse struct {
	CallInfo []*CallInfo `xml:"CallInfo,omitempty" json:"CallInfo,omitempty" yaml:"CallInfo,omitempty"`
}

// Device was auto-generated from WSDL.
type Device struct {
	Hw   string `xml:"hw" json:"hw" yaml:"hw"`
	Text string `xml:"text" json:"text" yaml:"text"`
}

// Devices was auto-generated from WSDL.
type Devices struct {
	Session int     `xml:"session" json:"session" yaml:"session"`
	User    *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// DevicesResponse was auto-generated from WSDL.
type DevicesResponse struct {
	Device []*Device `xml:"Device,omitempty" json:"Device,omitempty" yaml:"Device,omitempty"`
}

// Echo was auto-generated from WSDL.
type Echo struct {
	Session int `xml:"session" json:"session" yaml:"session"`
	Key     int `xml:"key" json:"key" yaml:"key"`
}

// EchoResponse was auto-generated from WSDL.
type EchoResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
}

// End was auto-generated from WSDL.
type End struct {
	Session int `xml:"session" json:"session" yaml:"session"`
}

// EndResponse was auto-generated from WSDL.
type EndResponse struct {
}

// FindUser was auto-generated from WSDL.
type FindUser struct {
	V501   *string `xml:"v501,omitempty" json:"v501,omitempty" yaml:"v501,omitempty"`
	V700   *string `xml:"v700,omitempty" json:"v700,omitempty" yaml:"v700,omitempty"`
	V800   *string `xml:"v800,omitempty" json:"v800,omitempty" yaml:"v800,omitempty"`
	Vx1100 *string `xml:"vx1100,omitempty" json:"vx1100,omitempty" yaml:"vx1100,omitempty"`
	Cn     *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	H323   *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	E164   *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	Count  int     `xml:"count" json:"count" yaml:"count"`
	Next   int     `xml:"next" json:"next" yaml:"next"`
	Nohide bool    `xml:"nohide" json:"nohide" yaml:"nohide"`
}

// FindUserResponse was auto-generated from WSDL.
type FindUserResponse struct {
	Return *UserInfoArray `xml:"return" json:"return" yaml:"return"`
}

// Group was auto-generated from WSDL.
type Group struct {
	Group  string `xml:"group" json:"group" yaml:"group"`
	Active bool   `xml:"active" json:"active" yaml:"active"`
}

// Info was auto-generated from WSDL.
type Info struct {
	Type *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
	Vali *int    `xml:"vali,omitempty" json:"vali,omitempty" yaml:"vali,omitempty"`
	Vals *string `xml:"vals,omitempty" json:"vals,omitempty" yaml:"vals,omitempty"`
}

// Initialize was auto-generated from WSDL.
type Initialize struct {
	User   *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Appl   *string `xml:"appl,omitempty" json:"appl,omitempty" yaml:"appl,omitempty"`
	V      bool    `xml:"v" json:"v" yaml:"v"`
	V501   bool    `xml:"v501" json:"v501" yaml:"v501"`
	V700   bool    `xml:"v700" json:"v700" yaml:"v700"`
	V800   bool    `xml:"v800" json:"v800" yaml:"v800"`
	Vx1100 bool    `xml:"vx1100" json:"vx1100" yaml:"vx1100"`
}

// InitializeResponse was auto-generated from WSDL.
type InitializeResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
	Key    int `xml:"key" json:"key" yaml:"key"`
}

// License was auto-generated from WSDL.
type License struct {
	Session int     `xml:"session" json:"session" yaml:"session"`
	Name    *string `xml:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty"`
}

// LicenseResponse was auto-generated from WSDL.
type LicenseResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// LocationUrl was auto-generated from WSDL.
type LocationUrl struct {
	V501   *string `xml:"v501,omitempty" json:"v501,omitempty" yaml:"v501,omitempty"`
	V700   *string `xml:"v700,omitempty" json:"v700,omitempty" yaml:"v700,omitempty"`
	V800   *string `xml:"v800,omitempty" json:"v800,omitempty" yaml:"v800,omitempty"`
	Vx1100 *string `xml:"vx1100,omitempty" json:"vx1100,omitempty" yaml:"vx1100,omitempty"`
	Loc    *string `xml:"loc,omitempty" json:"loc,omitempty" yaml:"loc,omitempty"`
	Tls    bool    `xml:"tls" json:"tls" yaml:"tls"`
}

// LocationUrlResponse was auto-generated from WSDL.
type LocationUrlResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// No was auto-generated from WSDL.
type No struct {
	Type string `xml:"type" json:"type" yaml:"type"`
	Cn   string `xml:"cn" json:"cn" yaml:"cn"`
	Dn   string `xml:"dn" json:"dn" yaml:"dn"`
	E164 string `xml:"e164" json:"e164" yaml:"e164"`
	H323 string `xml:"h323" json:"h323" yaml:"h323"`
}

// Poll was auto-generated from WSDL.
type Poll struct {
	Session int `xml:"session" json:"session" yaml:"session"`
}

// PollResponse was auto-generated from WSDL.
type PollResponse struct {
	Return *AnyInfo `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Presence was auto-generated from WSDL.
type Presence struct {
	Status   *PresenceStatus   `xml:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty"`
	Activity *PresenceActivity `xml:"activity,omitempty" json:"activity,omitempty" yaml:"activity,omitempty"`
	Note     *string           `xml:"note,omitempty" json:"note,omitempty" yaml:"note,omitempty"`
}

// RegInfo was auto-generated from WSDL.
type RegInfo struct {
	Active *bool   `xml:"active,omitempty" json:"active,omitempty" yaml:"active,omitempty"`
	User   *int    `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Reg    *int    `xml:"reg,omitempty" json:"reg,omitempty" yaml:"reg,omitempty"`
	Hw     *string `xml:"hw,omitempty" json:"hw,omitempty" yaml:"hw,omitempty"`
	Soap   *string `xml:"soap,omitempty" json:"soap,omitempty" yaml:"soap,omitempty"`
	Info   []*Info `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// SetPresence was auto-generated from WSDL.
type SetPresence struct {
	Presence *Presence `xml:"presence,omitempty" json:"presence,omitempty" yaml:"presence,omitempty"`
	Im       bool      `xml:"im" json:"im" yaml:"im"`
	Contact  *string   `xml:"contact,omitempty" json:"contact,omitempty" yaml:"contact,omitempty"`
	Guid     *string   `xml:"guid,omitempty" json:"guid,omitempty" yaml:"guid,omitempty"`
	H323     *string   `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
}

// SetPresenceResponse was auto-generated from WSDL.
type SetPresenceResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
}

// UserCall was auto-generated from WSDL.
type UserCall struct {
	User    int     `xml:"user" json:"user" yaml:"user"`
	Cn      *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	E164    *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323    *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Reg     int     `xml:"reg" json:"reg" yaml:"reg"`
	Info    []*Info `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
	Rc      int     `xml:"rc" json:"rc" yaml:"rc"`
	Srce164 *string `xml:"srce164,omitempty" json:"srce164,omitempty" yaml:"srce164,omitempty"`
}

// UserCallResponse was auto-generated from WSDL.
type UserCallResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
}

// UserClear was auto-generated from WSDL.
type UserClear struct {
	Call  int     `xml:"call" json:"call" yaml:"call"`
	Cause int     `xml:"cause" json:"cause" yaml:"cause"`
	Info  []*Info `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// UserClearResponse was auto-generated from WSDL.
type UserClearResponse struct {
}

// UserConnect was auto-generated from WSDL.
type UserConnect struct {
	Call int `xml:"call" json:"call" yaml:"call"`
}

// UserConnectResponse was auto-generated from WSDL.
type UserConnectResponse struct {
}

// UserCtComplete was auto-generated from WSDL.
type UserCtComplete struct {
	Call int     `xml:"call" json:"call" yaml:"call"`
	E164 *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
}

// UserCtCompleteResponse was auto-generated from WSDL.
type UserCtCompleteResponse struct {
}

// UserDTMF was auto-generated from WSDL.
type UserDTMF struct {
	Call int     `xml:"call" json:"call" yaml:"call"`
	Recv bool    `xml:"recv" json:"recv" yaml:"recv"`
	Dtmf *string `xml:"dtmf,omitempty" json:"dtmf,omitempty" yaml:"dtmf,omitempty"`
}

// UserDTMFResponse was auto-generated from WSDL.
type UserDTMFResponse struct {
}

// UserEnd was auto-generated from WSDL.
type UserEnd struct {
	User int `xml:"user" json:"user" yaml:"user"`
}

// UserEndResponse was auto-generated from WSDL.
type UserEndResponse struct {
}

// UserFindDestination was auto-generated from WSDL.
type UserFindDestination struct {
	User int     `xml:"user" json:"user" yaml:"user"`
	E164 *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
}

// UserFindDestinationResponse was auto-generated from WSDL.
type UserFindDestinationResponse struct {
	Return bool      `xml:"return" json:"return" yaml:"return"`
	User1  *UserInfo `xml:"user1,omitempty" json:"user1,omitempty" yaml:"user1,omitempty"`
}

// UserHold was auto-generated from WSDL.
type UserHold struct {
	Call   int  `xml:"call" json:"call" yaml:"call"`
	Remote bool `xml:"remote" json:"remote" yaml:"remote"`
}

// UserHoldResponse was auto-generated from WSDL.
type UserHoldResponse struct {
}

// UserInfo was auto-generated from WSDL.
type UserInfo struct {
	Active    bool      `xml:"active" json:"active" yaml:"active"`
	Cn        string    `xml:"cn" json:"cn" yaml:"cn"`
	Dn        string    `xml:"dn" json:"dn" yaml:"dn"`
	Type      string    `xml:"type" json:"type" yaml:"type"`
	E164      string    `xml:"e164" json:"e164" yaml:"e164"`
	H323      string    `xml:"h323" json:"h323" yaml:"h323"`
	State     int       `xml:"state" json:"state" yaml:"state"`
	Channel   int       `xml:"channel" json:"channel" yaml:"channel"`
	Alert     int       `xml:"alert" json:"alert" yaml:"alert"`
	Guid      string    `xml:"guid" json:"guid" yaml:"guid"`
	Groups    []*Group  `xml:"groups,omitempty" json:"groups,omitempty" yaml:"groups,omitempty"`
	Loc       *string   `xml:"loc,omitempty" json:"loc,omitempty" yaml:"loc,omitempty"`
	Node      *string   `xml:"node,omitempty" json:"node,omitempty" yaml:"node,omitempty"`
	Nodenum   *string   `xml:"nodenum,omitempty" json:"nodenum,omitempty" yaml:"nodenum,omitempty"`
	Object    *string   `xml:"object,omitempty" json:"object,omitempty" yaml:"object,omitempty"`
	Cfg       *bool     `xml:"cfg,omitempty" json:"cfg,omitempty" yaml:"cfg,omitempty"`
	Presence  *Presence `xml:"presence" json:"presence" yaml:"presence"`
	Domain    *string   `xml:"domain,omitempty" json:"domain,omitempty" yaml:"domain,omitempty"`
	H323email *bool     `xml:"h323email,omitempty" json:"h323email,omitempty" yaml:"h323email,omitempty"`
	Email     []*string `xml:"email,omitempty" json:"email,omitempty" yaml:"email,omitempty"`
	Fake      *string   `xml:"fake,omitempty" json:"fake,omitempty" yaml:"fake,omitempty"`
	Info      []*Info   `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// UserInfoArray was auto-generated from WSDL.
type UserInfoArray struct {
	User []*UserInfo `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
}

// UserInfoTx was auto-generated from WSDL.
type UserInfoTx struct {
	Call int     `xml:"call" json:"call" yaml:"call"`
	Recv bool    `xml:"recv" json:"recv" yaml:"recv"`
	Cdpn *string `xml:"cdpn,omitempty" json:"cdpn,omitempty" yaml:"cdpn,omitempty"`
	Key  *string `xml:"key,omitempty" json:"key,omitempty" yaml:"key,omitempty"`
	Dsp  *string `xml:"dsp,omitempty" json:"dsp,omitempty" yaml:"dsp,omitempty"`
}

// UserInfoTxResponse was auto-generated from WSDL.
type UserInfoTxResponse struct {
}

// UserInitialize was auto-generated from WSDL.
type UserInitialize struct {
	Session int     `xml:"session" json:"session" yaml:"session"`
	User    *string `xml:"user,omitempty" json:"user,omitempty" yaml:"user,omitempty"`
	Xfer    bool    `xml:"xfer" json:"xfer" yaml:"xfer"`
	Disc    bool    `xml:"disc" json:"disc" yaml:"disc"`
	Hw      *string `xml:"hw,omitempty" json:"hw,omitempty" yaml:"hw,omitempty"`
}

// UserInitializeResponse was auto-generated from WSDL.
type UserInitializeResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
}

// UserLicense was auto-generated from WSDL.
type UserLicense struct {
	Guid *string `xml:"guid,omitempty" json:"guid,omitempty" yaml:"guid,omitempty"`
	Type *string `xml:"type,omitempty" json:"type,omitempty" yaml:"type,omitempty"`
}

// UserLicenseResponse was auto-generated from WSDL.
type UserLicenseResponse struct {
	Return bool `xml:"return" json:"return" yaml:"return"`
}

// UserLocalNum was auto-generated from WSDL.
type UserLocalNum struct {
	User int     `xml:"user" json:"user" yaml:"user"`
	Num  *string `xml:"num,omitempty" json:"num,omitempty" yaml:"num,omitempty"`
}

// UserLocalNumResponse was auto-generated from WSDL.
type UserLocalNumResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// UserMediaTransfer was auto-generated from WSDL.
type UserMediaTransfer struct {
	A    int  `xml:"a" json:"a" yaml:"a"`
	B    int  `xml:"b" json:"b" yaml:"b"`
	User bool `xml:"user" json:"user" yaml:"user"`
	Peer bool `xml:"peer" json:"peer" yaml:"peer"`
}

// UserMediaTransferResponse was auto-generated from WSDL.
type UserMediaTransferResponse struct {
}

// UserMessage was auto-generated from WSDL.
type UserMessage struct {
	User     int     `xml:"user" json:"user" yaml:"user"`
	E164     *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323     *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Msg      *string `xml:"msg,omitempty" json:"msg,omitempty" yaml:"msg,omitempty"`
	Src_e164 *string `xml:"src_e164,omitempty" json:"src_e164,omitempty" yaml:"src_e164,omitempty"`
	Src_h323 *string `xml:"src_h323,omitempty" json:"src_h323,omitempty" yaml:"src_h323,omitempty"`
}

// UserMessageResponse was auto-generated from WSDL.
type UserMessageResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
}

// UserPark was auto-generated from WSDL.
type UserPark struct {
	Call     int     `xml:"call" json:"call" yaml:"call"`
	Cn       *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	Position int     `xml:"position" json:"position" yaml:"position"`
}

// UserParkResponse was auto-generated from WSDL.
type UserParkResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
}

// UserPickup was auto-generated from WSDL.
type UserPickup struct {
	User  int     `xml:"user" json:"user" yaml:"user"`
	Cn    *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	Call  int     `xml:"call" json:"call" yaml:"call"`
	Group *string `xml:"group,omitempty" json:"group,omitempty" yaml:"group,omitempty"`
	Reg   int     `xml:"reg" json:"reg" yaml:"reg"`
	Info  []*Info `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// UserPickupResponse was auto-generated from WSDL.
type UserPickupResponse struct {
	Return int `xml:"return" json:"return" yaml:"return"`
}

// UserRc was auto-generated from WSDL.
type UserRc struct {
	Call int `xml:"call" json:"call" yaml:"call"`
	Rc   int `xml:"rc" json:"rc" yaml:"rc"`
}

// UserRcResponse was auto-generated from WSDL.
type UserRcResponse struct {
}

// UserRedirect was auto-generated from WSDL.
type UserRedirect struct {
	Call int     `xml:"call" json:"call" yaml:"call"`
	Cn   *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	E164 *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Info []*Info `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
	Rc   int     `xml:"rc" json:"rc" yaml:"rc"`
}

// UserRedirectResponse was auto-generated from WSDL.
type UserRedirectResponse struct {
	Return bool `xml:"return" json:"return" yaml:"return"`
}

// UserReroute was auto-generated from WSDL.
type UserReroute struct {
	Call int     `xml:"call" json:"call" yaml:"call"`
	Cn   *string `xml:"cn,omitempty" json:"cn,omitempty" yaml:"cn,omitempty"`
	E164 *string `xml:"e164,omitempty" json:"e164,omitempty" yaml:"e164,omitempty"`
	H323 *string `xml:"h323,omitempty" json:"h323,omitempty" yaml:"h323,omitempty"`
	Info []*Info `xml:"info,omitempty" json:"info,omitempty" yaml:"info,omitempty"`
}

// UserRerouteResponse was auto-generated from WSDL.
type UserRerouteResponse struct {
	Return bool `xml:"return" json:"return" yaml:"return"`
}

// UserRetrieve was auto-generated from WSDL.
type UserRetrieve struct {
	Call int `xml:"call" json:"call" yaml:"call"`
}

// UserRetrieveResponse was auto-generated from WSDL.
type UserRetrieveResponse struct {
}

// UserTransfer was auto-generated from WSDL.
type UserTransfer struct {
	A int `xml:"a" json:"a" yaml:"a"`
	B int `xml:"b" json:"b" yaml:"b"`
}

// UserTransferResponse was auto-generated from WSDL.
type UserTransferResponse struct {
}

// UserUUI was auto-generated from WSDL.
type UserUUI struct {
	Call int     `xml:"call" json:"call" yaml:"call"`
	Recv bool    `xml:"recv" json:"recv" yaml:"recv"`
	Uui  *string `xml:"uui,omitempty" json:"uui,omitempty" yaml:"uui,omitempty"`
}

// UserUUIResponse was auto-generated from WSDL.
type UserUUIResponse struct {
}

// Version was auto-generated from WSDL.
type Version struct {
}

// VersionResponse was auto-generated from WSDL.
type VersionResponse struct {
	VersionResult   int     `xml:"VersionResult" json:"VersionResult" yaml:"VersionResult"`
	GatekeeperID    *string `xml:"GatekeeperID,omitempty" json:"GatekeeperID,omitempty" yaml:"GatekeeperID,omitempty"`
	Location        *string `xml:"Location,omitempty" json:"Location,omitempty" yaml:"Location,omitempty"`
	FirmwareVersion *string `xml:"FirmwareVersion,omitempty" json:"FirmwareVersion,omitempty" yaml:"FirmwareVersion,omitempty"`
	SerialNumber    *string `xml:"SerialNumber,omitempty" json:"SerialNumber,omitempty" yaml:"SerialNumber,omitempty"`
}

// Operation wrapper for Admin.
// OperationAdminRequest was auto-generated from WSDL.
type OperationAdminRequest struct {
	Admin *Admin `xml:"Admin,omitempty" json:"Admin,omitempty" yaml:"Admin,omitempty"`
}

// Operation wrapper for Admin.
// OperationAdminResponse was auto-generated from WSDL.
type OperationAdminResponse struct {
	AdminResponse *AdminResponse `xml:"AdminResponse,omitempty" json:"AdminResponse,omitempty" yaml:"AdminResponse,omitempty"`
}

// Operation wrapper for Calls.
// OperationCallsRequest was auto-generated from WSDL.
type OperationCallsRequest struct {
	Calls *Calls `xml:"Calls,omitempty" json:"Calls,omitempty" yaml:"Calls,omitempty"`
}

// Operation wrapper for Calls.
// OperationCallsResponse was auto-generated from WSDL.
type OperationCallsResponse struct {
	CallsResponse *CallsResponse `xml:"CallsResponse,omitempty" json:"CallsResponse,omitempty" yaml:"CallsResponse,omitempty"`
}

// Operation wrapper for Devices.
// OperationDevicesRequest was auto-generated from WSDL.
type OperationDevicesRequest struct {
	Devices *Devices `xml:"Devices,omitempty" json:"Devices,omitempty" yaml:"Devices,omitempty"`
}

// Operation wrapper for Devices.
// OperationDevicesResponse was auto-generated from WSDL.
type OperationDevicesResponse struct {
	DevicesResponse *DevicesResponse `xml:"DevicesResponse,omitempty" json:"DevicesResponse,omitempty" yaml:"DevicesResponse,omitempty"`
}

// Operation wrapper for Echo.
// OperationEchoRequest was auto-generated from WSDL.
type OperationEchoRequest struct {
	Echo *Echo `xml:"Echo,omitempty" json:"Echo,omitempty" yaml:"Echo,omitempty"`
}

// Operation wrapper for Echo.
// OperationEchoResponse was auto-generated from WSDL.
type OperationEchoResponse struct {
	EchoResponse *EchoResponse `xml:"EchoResponse,omitempty" json:"EchoResponse,omitempty" yaml:"EchoResponse,omitempty"`
}

// Operation wrapper for End.
// OperationEndRequest was auto-generated from WSDL.
type OperationEndRequest struct {
	End *End `xml:"End,omitempty" json:"End,omitempty" yaml:"End,omitempty"`
}

// Operation wrapper for End.
// OperationEndResponse was auto-generated from WSDL.
type OperationEndResponse struct {
	EndResponse *EndResponse `xml:"EndResponse,omitempty" json:"EndResponse,omitempty" yaml:"EndResponse,omitempty"`
}

// Operation wrapper for FindUser.
// OperationFindUserRequest was auto-generated from WSDL.
type OperationFindUserRequest struct {
	FindUser *FindUser `xml:"FindUser,omitempty" json:"FindUser,omitempty" yaml:"FindUser,omitempty"`
}

// Operation wrapper for FindUser.
// OperationFindUserResponse was auto-generated from WSDL.
type OperationFindUserResponse struct {
	FindUserResponse *FindUserResponse `xml:"FindUserResponse,omitempty" json:"FindUserResponse,omitempty" yaml:"FindUserResponse,omitempty"`
}

// Operation wrapper for Initialize.
// OperationInitializeRequest was auto-generated from WSDL.
type OperationInitializeRequest struct {
	Initialize *Initialize `xml:"Initialize,omitempty" json:"Initialize,omitempty" yaml:"Initialize,omitempty"`
}

// Operation wrapper for Initialize.
// OperationInitializeResponse was auto-generated from WSDL.
type OperationInitializeResponse struct {
	InitializeResponse *InitializeResponse `xml:"InitializeResponse,omitempty" json:"InitializeResponse,omitempty" yaml:"InitializeResponse,omitempty"`
}

// Operation wrapper for License.
// OperationLicenseRequest was auto-generated from WSDL.
type OperationLicenseRequest struct {
	License *License `xml:"License,omitempty" json:"License,omitempty" yaml:"License,omitempty"`
}

// Operation wrapper for License.
// OperationLicenseResponse was auto-generated from WSDL.
type OperationLicenseResponse struct {
	LicenseResponse *LicenseResponse `xml:"LicenseResponse,omitempty" json:"LicenseResponse,omitempty" yaml:"LicenseResponse,omitempty"`
}

// Operation wrapper for LocationUrl.
// OperationLocationUrlRequest was auto-generated from WSDL.
type OperationLocationUrlRequest struct {
	LocationUrl *LocationUrl `xml:"LocationUrl,omitempty" json:"LocationUrl,omitempty" yaml:"LocationUrl,omitempty"`
}

// Operation wrapper for LocationUrl.
// OperationLocationUrlResponse was auto-generated from WSDL.
type OperationLocationUrlResponse struct {
	LocationUrlResponse *LocationUrlResponse `xml:"LocationUrlResponse,omitempty" json:"LocationUrlResponse,omitempty" yaml:"LocationUrlResponse,omitempty"`
}

// Operation wrapper for Poll.
// OperationPollRequest was auto-generated from WSDL.
type OperationPollRequest struct {
	Poll *Poll `xml:"Poll,omitempty" json:"Poll,omitempty" yaml:"Poll,omitempty"`
}

// Operation wrapper for Poll.
// OperationPollResponse was auto-generated from WSDL.
type OperationPollResponse struct {
	PollResponse *PollResponse `xml:"PollResponse,omitempty" json:"PollResponse,omitempty" yaml:"PollResponse,omitempty"`
}

// Operation wrapper for SetPresence.
// OperationSetPresenceRequest was auto-generated from WSDL.
type OperationSetPresenceRequest struct {
	SetPresence *SetPresence `xml:"SetPresence,omitempty" json:"SetPresence,omitempty" yaml:"SetPresence,omitempty"`
}

// Operation wrapper for SetPresence.
// OperationSetPresenceResponse was auto-generated from WSDL.
type OperationSetPresenceResponse struct {
	SetPresenceResponse *SetPresenceResponse `xml:"SetPresenceResponse,omitempty" json:"SetPresenceResponse,omitempty" yaml:"SetPresenceResponse,omitempty"`
}

// Operation wrapper for UserCall.
// OperationUserCallRequest was auto-generated from WSDL.
type OperationUserCallRequest struct {
	UserCall *UserCall `xml:"UserCall,omitempty" json:"UserCall,omitempty" yaml:"UserCall,omitempty"`
}

// Operation wrapper for UserCall.
// OperationUserCallResponse was auto-generated from WSDL.
type OperationUserCallResponse struct {
	UserCallResponse *UserCallResponse `xml:"UserCallResponse,omitempty" json:"UserCallResponse,omitempty" yaml:"UserCallResponse,omitempty"`
}

// Operation wrapper for UserClear.
// OperationUserClearRequest was auto-generated from WSDL.
type OperationUserClearRequest struct {
	UserClear *UserClear `xml:"UserClear,omitempty" json:"UserClear,omitempty" yaml:"UserClear,omitempty"`
}

// Operation wrapper for UserClear.
// OperationUserClearResponse was auto-generated from WSDL.
type OperationUserClearResponse struct {
	UserClearResponse *UserClearResponse `xml:"UserClearResponse,omitempty" json:"UserClearResponse,omitempty" yaml:"UserClearResponse,omitempty"`
}

// Operation wrapper for UserConnect.
// OperationUserConnectRequest was auto-generated from WSDL.
type OperationUserConnectRequest struct {
	UserConnect *UserConnect `xml:"UserConnect,omitempty" json:"UserConnect,omitempty" yaml:"UserConnect,omitempty"`
}

// Operation wrapper for UserConnect.
// OperationUserConnectResponse was auto-generated from WSDL.
type OperationUserConnectResponse struct {
	UserConnectResponse *UserConnectResponse `xml:"UserConnectResponse,omitempty" json:"UserConnectResponse,omitempty" yaml:"UserConnectResponse,omitempty"`
}

// Operation wrapper for UserCtComplete.
// OperationUserCtCompleteRequest was auto-generated from WSDL.
type OperationUserCtCompleteRequest struct {
	UserCtComplete *UserCtComplete `xml:"UserCtComplete,omitempty" json:"UserCtComplete,omitempty" yaml:"UserCtComplete,omitempty"`
}

// Operation wrapper for UserCtComplete.
// OperationUserCtCompleteResponse was auto-generated from WSDL.
type OperationUserCtCompleteResponse struct {
	UserCtCompleteResponse *UserCtCompleteResponse `xml:"UserCtCompleteResponse,omitempty" json:"UserCtCompleteResponse,omitempty" yaml:"UserCtCompleteResponse,omitempty"`
}

// Operation wrapper for UserDTMF.
// OperationUserDTMFRequest was auto-generated from WSDL.
type OperationUserDTMFRequest struct {
	UserDTMF *UserDTMF `xml:"UserDTMF,omitempty" json:"UserDTMF,omitempty" yaml:"UserDTMF,omitempty"`
}

// Operation wrapper for UserDTMF.
// OperationUserDTMFResponse was auto-generated from WSDL.
type OperationUserDTMFResponse struct {
	UserDTMFResponse *UserDTMFResponse `xml:"UserDTMFResponse,omitempty" json:"UserDTMFResponse,omitempty" yaml:"UserDTMFResponse,omitempty"`
}

// Operation wrapper for UserEnd.
// OperationUserEndRequest was auto-generated from WSDL.
type OperationUserEndRequest struct {
	UserEnd *UserEnd `xml:"UserEnd,omitempty" json:"UserEnd,omitempty" yaml:"UserEnd,omitempty"`
}

// Operation wrapper for UserEnd.
// OperationUserEndResponse was auto-generated from WSDL.
type OperationUserEndResponse struct {
	UserEndResponse *UserEndResponse `xml:"UserEndResponse,omitempty" json:"UserEndResponse,omitempty" yaml:"UserEndResponse,omitempty"`
}

// Operation wrapper for UserFindDestination.
// OperationUserFindDestinationRequest was auto-generated from
// WSDL.
type OperationUserFindDestinationRequest struct {
	UserFindDestination *UserFindDestination `xml:"UserFindDestination,omitempty" json:"UserFindDestination,omitempty" yaml:"UserFindDestination,omitempty"`
}

// Operation wrapper for UserFindDestination.
// OperationUserFindDestinationResponse was auto-generated from
// WSDL.
type OperationUserFindDestinationResponse struct {
	UserFindDestinationResponse *UserFindDestinationResponse `xml:"UserFindDestinationResponse,omitempty" json:"UserFindDestinationResponse,omitempty" yaml:"UserFindDestinationResponse,omitempty"`
}

// Operation wrapper for UserHold.
// OperationUserHoldRequest was auto-generated from WSDL.
type OperationUserHoldRequest struct {
	UserHold *UserHold `xml:"UserHold,omitempty" json:"UserHold,omitempty" yaml:"UserHold,omitempty"`
}

// Operation wrapper for UserHold.
// OperationUserHoldResponse was auto-generated from WSDL.
type OperationUserHoldResponse struct {
	UserHoldResponse *UserHoldResponse `xml:"UserHoldResponse,omitempty" json:"UserHoldResponse,omitempty" yaml:"UserHoldResponse,omitempty"`
}

// Operation wrapper for UserInfo.
// OperationUserInfoRequest was auto-generated from WSDL.
type OperationUserInfoRequest struct {
	UserInfoTx *UserInfoTx `xml:"UserInfoTx,omitempty" json:"UserInfoTx,omitempty" yaml:"UserInfoTx,omitempty"`
}

// Operation wrapper for UserInfo.
// OperationUserInfoResponse was auto-generated from WSDL.
type OperationUserInfoResponse struct {
	UserInfoTxResponse *UserInfoTxResponse `xml:"UserInfoTxResponse,omitempty" json:"UserInfoTxResponse,omitempty" yaml:"UserInfoTxResponse,omitempty"`
}

// Operation wrapper for UserInitialize.
// OperationUserInitializeRequest was auto-generated from WSDL.
type OperationUserInitializeRequest struct {
	UserInitialize *UserInitialize `xml:"UserInitialize,omitempty" json:"UserInitialize,omitempty" yaml:"UserInitialize,omitempty"`
}

// Operation wrapper for UserInitialize.
// OperationUserInitializeResponse was auto-generated from WSDL.
type OperationUserInitializeResponse struct {
	UserInitializeResponse *UserInitializeResponse `xml:"UserInitializeResponse,omitempty" json:"UserInitializeResponse,omitempty" yaml:"UserInitializeResponse,omitempty"`
}

// Operation wrapper for UserLicense.
// OperationUserLicenseRequest was auto-generated from WSDL.
type OperationUserLicenseRequest struct {
	UserLicense *UserLicense `xml:"UserLicense,omitempty" json:"UserLicense,omitempty" yaml:"UserLicense,omitempty"`
}

// Operation wrapper for UserLicense.
// OperationUserLicenseResponse was auto-generated from WSDL.
type OperationUserLicenseResponse struct {
	UserLicenseResponse *UserLicenseResponse `xml:"UserLicenseResponse,omitempty" json:"UserLicenseResponse,omitempty" yaml:"UserLicenseResponse,omitempty"`
}

// Operation wrapper for UserLocalNum.
// OperationUserLocalNumRequest was auto-generated from WSDL.
type OperationUserLocalNumRequest struct {
	UserLocalNum *UserLocalNum `xml:"UserLocalNum,omitempty" json:"UserLocalNum,omitempty" yaml:"UserLocalNum,omitempty"`
}

// Operation wrapper for UserLocalNum.
// OperationUserLocalNumResponse was auto-generated from WSDL.
type OperationUserLocalNumResponse struct {
	UserLocalNumResponse *UserLocalNumResponse `xml:"UserLocalNumResponse,omitempty" json:"UserLocalNumResponse,omitempty" yaml:"UserLocalNumResponse,omitempty"`
}

// Operation wrapper for UserMediaTransfer.
// OperationUserMediaTransferRequest was auto-generated from WSDL.
type OperationUserMediaTransferRequest struct {
	UserMediaTransfer *UserMediaTransfer `xml:"UserMediaTransfer,omitempty" json:"UserMediaTransfer,omitempty" yaml:"UserMediaTransfer,omitempty"`
}

// Operation wrapper for UserMediaTransfer.
// OperationUserMediaTransferResponse was auto-generated from WSDL.
type OperationUserMediaTransferResponse struct {
	UserMediaTransferResponse *UserMediaTransferResponse `xml:"UserMediaTransferResponse,omitempty" json:"UserMediaTransferResponse,omitempty" yaml:"UserMediaTransferResponse,omitempty"`
}

// Operation wrapper for UserMessage.
// OperationUserMessageRequest was auto-generated from WSDL.
type OperationUserMessageRequest struct {
	UserMessage *UserMessage `xml:"UserMessage,omitempty" json:"UserMessage,omitempty" yaml:"UserMessage,omitempty"`
}

// Operation wrapper for UserMessage.
// OperationUserMessageResponse was auto-generated from WSDL.
type OperationUserMessageResponse struct {
	UserMessageResponse *UserMessageResponse `xml:"UserMessageResponse,omitempty" json:"UserMessageResponse,omitempty" yaml:"UserMessageResponse,omitempty"`
}

// Operation wrapper for UserPark.
// OperationUserParkRequest was auto-generated from WSDL.
type OperationUserParkRequest struct {
	UserPark *UserPark `xml:"UserPark,omitempty" json:"UserPark,omitempty" yaml:"UserPark,omitempty"`
}

// Operation wrapper for UserPark.
// OperationUserParkResponse was auto-generated from WSDL.
type OperationUserParkResponse struct {
	UserParkResponse *UserParkResponse `xml:"UserParkResponse,omitempty" json:"UserParkResponse,omitempty" yaml:"UserParkResponse,omitempty"`
}

// Operation wrapper for UserPickup.
// OperationUserPickupRequest was auto-generated from WSDL.
type OperationUserPickupRequest struct {
	UserPickup *UserPickup `xml:"UserPickup,omitempty" json:"UserPickup,omitempty" yaml:"UserPickup,omitempty"`
}

// Operation wrapper for UserPickup.
// OperationUserPickupResponse was auto-generated from WSDL.
type OperationUserPickupResponse struct {
	UserPickupResponse *UserPickupResponse `xml:"UserPickupResponse,omitempty" json:"UserPickupResponse,omitempty" yaml:"UserPickupResponse,omitempty"`
}

// Operation wrapper for UserRc.
// OperationUserRcRequest was auto-generated from WSDL.
type OperationUserRcRequest struct {
	UserRc *UserRc `xml:"UserRc,omitempty" json:"UserRc,omitempty" yaml:"UserRc,omitempty"`
}

// Operation wrapper for UserRc.
// OperationUserRcResponse was auto-generated from WSDL.
type OperationUserRcResponse struct {
	UserRcResponse *UserRcResponse `xml:"UserRcResponse,omitempty" json:"UserRcResponse,omitempty" yaml:"UserRcResponse,omitempty"`
}

// Operation wrapper for UserRedirect.
// OperationUserRedirectRequest was auto-generated from WSDL.
type OperationUserRedirectRequest struct {
	UserRedirect *UserRedirect `xml:"UserRedirect,omitempty" json:"UserRedirect,omitempty" yaml:"UserRedirect,omitempty"`
}

// Operation wrapper for UserRedirect.
// OperationUserRedirectResponse was auto-generated from WSDL.
type OperationUserRedirectResponse struct {
	UserRedirectResponse *UserRedirectResponse `xml:"UserRedirectResponse,omitempty" json:"UserRedirectResponse,omitempty" yaml:"UserRedirectResponse,omitempty"`
}

// Operation wrapper for UserReroute.
// OperationUserRerouteRequest was auto-generated from WSDL.
type OperationUserRerouteRequest struct {
	UserReroute *UserReroute `xml:"UserReroute,omitempty" json:"UserReroute,omitempty" yaml:"UserReroute,omitempty"`
}

// Operation wrapper for UserReroute.
// OperationUserRerouteResponse was auto-generated from WSDL.
type OperationUserRerouteResponse struct {
	UserRerouteResponse *UserRerouteResponse `xml:"UserRerouteResponse,omitempty" json:"UserRerouteResponse,omitempty" yaml:"UserRerouteResponse,omitempty"`
}

// Operation wrapper for UserRetrieve.
// OperationUserRetrieveRequest was auto-generated from WSDL.
type OperationUserRetrieveRequest struct {
	UserRetrieve *UserRetrieve `xml:"UserRetrieve,omitempty" json:"UserRetrieve,omitempty" yaml:"UserRetrieve,omitempty"`
}

// Operation wrapper for UserRetrieve.
// OperationUserRetrieveResponse was auto-generated from WSDL.
type OperationUserRetrieveResponse struct {
	UserRetrieveResponse *UserRetrieveResponse `xml:"UserRetrieveResponse,omitempty" json:"UserRetrieveResponse,omitempty" yaml:"UserRetrieveResponse,omitempty"`
}

// Operation wrapper for UserTransfer.
// OperationUserTransferRequest was auto-generated from WSDL.
type OperationUserTransferRequest struct {
	UserTransfer *UserTransfer `xml:"UserTransfer,omitempty" json:"UserTransfer,omitempty" yaml:"UserTransfer,omitempty"`
}

// Operation wrapper for UserTransfer.
// OperationUserTransferResponse was auto-generated from WSDL.
type OperationUserTransferResponse struct {
	UserTransferResponse *UserTransferResponse `xml:"UserTransferResponse,omitempty" json:"UserTransferResponse,omitempty" yaml:"UserTransferResponse,omitempty"`
}

// Operation wrapper for UserUUI.
// OperationUserUUIRequest was auto-generated from WSDL.
type OperationUserUUIRequest struct {
	UserUUI *UserUUI `xml:"UserUUI,omitempty" json:"UserUUI,omitempty" yaml:"UserUUI,omitempty"`
}

// Operation wrapper for UserUUI.
// OperationUserUUIResponse was auto-generated from WSDL.
type OperationUserUUIResponse struct {
	UserUUIResponse *UserUUIResponse `xml:"UserUUIResponse,omitempty" json:"UserUUIResponse,omitempty" yaml:"UserUUIResponse,omitempty"`
}

// Operation wrapper for Version.
// OperationVersionRequest was auto-generated from WSDL.
type OperationVersionRequest struct {
	Version *Version `xml:"Version,omitempty" json:"Version,omitempty" yaml:"Version,omitempty"`
}

// Operation wrapper for Version.
// OperationVersionResponse was auto-generated from WSDL.
type OperationVersionResponse struct {
	VersionResponse *VersionResponse `xml:"VersionResponse,omitempty" json:"VersionResponse,omitempty" yaml:"VersionResponse,omitempty"`
}

// pbxPortType implements the PbxPortType interface.
type pbxPortType struct {
	cli *soap.Client
}

// Admin was auto-generated from WSDL.
func (p *pbxPortType) Admin(Admin *Admin) (*AdminResponse, error) {
	α := struct {
		OperationAdminRequest `xml:"inno:Admin"`
	}{
		OperationAdminRequest{
			Admin,
		},
	}

	γ := struct {
		OperationAdminResponse `xml:"AdminResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/Admin", α, &γ); err != nil {
		return nil, err
	}
	return γ.AdminResponse, nil
}

// Calls was auto-generated from WSDL.
func (p *pbxPortType) Calls(Calls *Calls) (*CallsResponse, error) {
	α := struct {
		OperationCallsRequest `xml:"inno:Calls"`
	}{
		OperationCallsRequest{
			Calls,
		},
	}

	γ := struct {
		OperationCallsResponse `xml:"CallsResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/Calls", α, &γ); err != nil {
		return nil, err
	}
	return γ.CallsResponse, nil
}

// Devices was auto-generated from WSDL.
func (p *pbxPortType) Devices(Devices *Devices) (*DevicesResponse, error) {
	α := struct {
		OperationDevicesRequest `xml:"inno:Devices"`
	}{
		OperationDevicesRequest{
			Devices,
		},
	}

	γ := struct {
		OperationDevicesResponse `xml:"DevicesResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/Devices", α, &γ); err != nil {
		return nil, err
	}
	return γ.DevicesResponse, nil
}

// Echo was auto-generated from WSDL.
func (p *pbxPortType) Echo(Echo *Echo) (*EchoResponse, error) {
	α := struct {
		OperationEchoRequest `xml:"inno:Echo"`
	}{
		OperationEchoRequest{
			Echo,
		},
	}

	γ := struct {
		OperationEchoResponse `xml:"EchoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/Echo", α, &γ); err != nil {
		return nil, err
	}
	return γ.EchoResponse, nil
}

// End was auto-generated from WSDL.
func (p *pbxPortType) End(End *End) (*EndResponse, error) {
	α := struct {
		OperationEndRequest `xml:"inno:End"`
	}{
		OperationEndRequest{
			End,
		},
	}

	γ := struct {
		OperationEndResponse `xml:"EndResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/End", α, &γ); err != nil {
		return nil, err
	}
	return γ.EndResponse, nil
}

// FindUser was auto-generated from WSDL.
func (p *pbxPortType) FindUser(FindUser *FindUser) (*FindUserResponse, error) {
	α := struct {
		OperationFindUserRequest `xml:"inno:FindUser"`
	}{
		OperationFindUserRequest{
			FindUser,
		},
	}

	γ := struct {
		OperationFindUserResponse `xml:"FindUserResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/FindUser", α, &γ); err != nil {
		return nil, err
	}
	return γ.FindUserResponse, nil
}

// Initialize was auto-generated from WSDL.
func (p *pbxPortType) Initialize(Initialize *Initialize) (*InitializeResponse, error) {
	α := struct {
		OperationInitializeRequest `xml:"inno:Initialize"`
	}{
		OperationInitializeRequest{
			Initialize,
		},
	}

	γ := struct {
		OperationInitializeResponse `xml:"InitializeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/Initialize", α, &γ); err != nil {
		return nil, err
	}
	return γ.InitializeResponse, nil
}

// License was auto-generated from WSDL.
func (p *pbxPortType) License(License *License) (*LicenseResponse, error) {
	α := struct {
		OperationLicenseRequest `xml:"inno:License"`
	}{
		OperationLicenseRequest{
			License,
		},
	}

	γ := struct {
		OperationLicenseResponse `xml:"LicenseResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/License", α, &γ); err != nil {
		return nil, err
	}
	return γ.LicenseResponse, nil
}

// LocationUrl was auto-generated from WSDL.
func (p *pbxPortType) LocationUrl(LocationUrl *LocationUrl) (*LocationUrlResponse, error) {
	α := struct {
		OperationLocationUrlRequest `xml:"inno:LocationUrl"`
	}{
		OperationLocationUrlRequest{
			LocationUrl,
		},
	}

	γ := struct {
		OperationLocationUrlResponse `xml:"LocationUrlResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/LocationUrl", α, &γ); err != nil {
		return nil, err
	}
	return γ.LocationUrlResponse, nil
}

// Poll was auto-generated from WSDL.
func (p *pbxPortType) Poll(Poll *Poll) (*PollResponse, error) {
	α := struct {
		OperationPollRequest `xml:"inno:Poll"`
	}{
		OperationPollRequest{
			Poll,
		},
	}

	γ := struct {
		OperationPollResponse `xml:"PollResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/Poll", α, &γ); err != nil {
		return nil, err
	}
	return γ.PollResponse, nil
}

// SetPresence was auto-generated from WSDL.
func (p *pbxPortType) SetPresence(SetPresence *SetPresence) (*SetPresenceResponse, error) {
	α := struct {
		OperationSetPresenceRequest `xml:"inno:SetPresence"`
	}{
		OperationSetPresenceRequest{
			SetPresence,
		},
	}

	γ := struct {
		OperationSetPresenceResponse `xml:"SetPresenceResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/SetPresence", α, &γ); err != nil {
		return nil, err
	}
	return γ.SetPresenceResponse, nil
}

// UserCall was auto-generated from WSDL.
func (p *pbxPortType) UserCall(UserCall *UserCall) (*UserCallResponse, error) {
	α := struct {
		OperationUserCallRequest `xml:"inno:UserCall"`
	}{
		OperationUserCallRequest{
			UserCall,
		},
	}

	γ := struct {
		OperationUserCallResponse `xml:"UserCallResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserCall", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserCallResponse, nil
}

// UserClear was auto-generated from WSDL.
func (p *pbxPortType) UserClear(UserClear *UserClear) (*UserClearResponse, error) {
	α := struct {
		OperationUserClearRequest `xml:"inno:UserClear"`
	}{
		OperationUserClearRequest{
			UserClear,
		},
	}

	γ := struct {
		OperationUserClearResponse `xml:"UserClearResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserClear", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserClearResponse, nil
}

// UserConnect was auto-generated from WSDL.
func (p *pbxPortType) UserConnect(UserConnect *UserConnect) (*UserConnectResponse, error) {
	α := struct {
		OperationUserConnectRequest `xml:"inno:UserConnect"`
	}{
		OperationUserConnectRequest{
			UserConnect,
		},
	}

	γ := struct {
		OperationUserConnectResponse `xml:"UserConnectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserConnect", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserConnectResponse, nil
}

// UserCtComplete was auto-generated from WSDL.
func (p *pbxPortType) UserCtComplete(UserCtComplete *UserCtComplete) (*UserCtCompleteResponse, error) {
	α := struct {
		OperationUserCtCompleteRequest `xml:"inno:UserCtComplete"`
	}{
		OperationUserCtCompleteRequest{
			UserCtComplete,
		},
	}

	γ := struct {
		OperationUserCtCompleteResponse `xml:"UserCtCompleteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserCtComplete", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserCtCompleteResponse, nil
}

// UserDTMF was auto-generated from WSDL.
func (p *pbxPortType) UserDTMF(UserDTMF *UserDTMF) (*UserDTMFResponse, error) {
	α := struct {
		OperationUserDTMFRequest `xml:"inno:UserDTMF"`
	}{
		OperationUserDTMFRequest{
			UserDTMF,
		},
	}

	γ := struct {
		OperationUserDTMFResponse `xml:"UserDTMFResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserDTMF", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserDTMFResponse, nil
}

// UserEnd was auto-generated from WSDL.
func (p *pbxPortType) UserEnd(UserEnd *UserEnd) (*UserEndResponse, error) {
	α := struct {
		OperationUserEndRequest `xml:"inno:UserEnd"`
	}{
		OperationUserEndRequest{
			UserEnd,
		},
	}

	γ := struct {
		OperationUserEndResponse `xml:"UserEndResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserEnd", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserEndResponse, nil
}

// UserFindDestination was auto-generated from WSDL.
func (p *pbxPortType) UserFindDestination(UserFindDestination *UserFindDestination) (*UserFindDestinationResponse, error) {
	α := struct {
		OperationUserFindDestinationRequest `xml:"inno:UserFindDestination"`
	}{
		OperationUserFindDestinationRequest{
			UserFindDestination,
		},
	}

	γ := struct {
		OperationUserFindDestinationResponse `xml:"UserFindDestinationResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserFindDestination", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserFindDestinationResponse, nil
}

// UserHold was auto-generated from WSDL.
func (p *pbxPortType) UserHold(UserHold *UserHold) (*UserHoldResponse, error) {
	α := struct {
		OperationUserHoldRequest `xml:"inno:UserHold"`
	}{
		OperationUserHoldRequest{
			UserHold,
		},
	}

	γ := struct {
		OperationUserHoldResponse `xml:"UserHoldResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserHold", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserHoldResponse, nil
}

// UserInfo was auto-generated from WSDL.
func (p *pbxPortType) UserInfo(UserInfoTx *UserInfoTx) (*UserInfoTxResponse, error) {
	α := struct {
		OperationUserInfoRequest `xml:"inno:UserInfo"`
	}{
		OperationUserInfoRequest{
			UserInfoTx,
		},
	}

	γ := struct {
		OperationUserInfoResponse `xml:"UserInfoResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserInfo", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserInfoTxResponse, nil
}

// UserInitialize was auto-generated from WSDL.
func (p *pbxPortType) UserInitialize(UserInitialize *UserInitialize) (*UserInitializeResponse, error) {
	α := struct {
		OperationUserInitializeRequest `xml:"inno:UserInitialize"`
	}{
		OperationUserInitializeRequest{
			UserInitialize,
		},
	}

	γ := struct {
		OperationUserInitializeResponse `xml:"UserInitializeResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserInitialize", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserInitializeResponse, nil
}

// UserLicense was auto-generated from WSDL.
func (p *pbxPortType) UserLicense(UserLicense *UserLicense) (*UserLicenseResponse, error) {
	α := struct {
		OperationUserLicenseRequest `xml:"inno:UserLicense"`
	}{
		OperationUserLicenseRequest{
			UserLicense,
		},
	}

	γ := struct {
		OperationUserLicenseResponse `xml:"UserLicenseResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserLicense", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserLicenseResponse, nil
}

// UserLocalNum was auto-generated from WSDL.
func (p *pbxPortType) UserLocalNum(UserLocalNum *UserLocalNum) (*UserLocalNumResponse, error) {
	α := struct {
		OperationUserLocalNumRequest `xml:"inno:UserLocalNum"`
	}{
		OperationUserLocalNumRequest{
			UserLocalNum,
		},
	}

	γ := struct {
		OperationUserLocalNumResponse `xml:"UserLocalNumResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserLocalNum", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserLocalNumResponse, nil
}

// UserMediaTransfer was auto-generated from WSDL.
func (p *pbxPortType) UserMediaTransfer(UserMediaTransfer *UserMediaTransfer) (*UserMediaTransferResponse, error) {
	α := struct {
		OperationUserMediaTransferRequest `xml:"inno:UserMediaTransfer"`
	}{
		OperationUserMediaTransferRequest{
			UserMediaTransfer,
		},
	}

	γ := struct {
		OperationUserMediaTransferResponse `xml:"UserMediaTransferResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserMediaTransfer", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserMediaTransferResponse, nil
}

// UserMessage was auto-generated from WSDL.
func (p *pbxPortType) UserMessage(UserMessage *UserMessage) (*UserMessageResponse, error) {
	α := struct {
		OperationUserMessageRequest `xml:"inno:UserMessage"`
	}{
		OperationUserMessageRequest{
			UserMessage,
		},
	}

	γ := struct {
		OperationUserMessageResponse `xml:"UserMessageResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserMessage", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserMessageResponse, nil
}

// UserPark was auto-generated from WSDL.
func (p *pbxPortType) UserPark(UserPark *UserPark) (*UserParkResponse, error) {
	α := struct {
		OperationUserParkRequest `xml:"inno:UserPark"`
	}{
		OperationUserParkRequest{
			UserPark,
		},
	}

	γ := struct {
		OperationUserParkResponse `xml:"UserParkResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserPark", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserParkResponse, nil
}

// UserPickup was auto-generated from WSDL.
func (p *pbxPortType) UserPickup(UserPickup *UserPickup) (*UserPickupResponse, error) {
	α := struct {
		OperationUserPickupRequest `xml:"inno:UserPickup"`
	}{
		OperationUserPickupRequest{
			UserPickup,
		},
	}

	γ := struct {
		OperationUserPickupResponse `xml:"UserPickupResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserPickup", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserPickupResponse, nil
}

// UserRc was auto-generated from WSDL.
func (p *pbxPortType) UserRc(UserRc *UserRc) (*UserRcResponse, error) {
	α := struct {
		OperationUserRcRequest `xml:"inno:UserRc"`
	}{
		OperationUserRcRequest{
			UserRc,
		},
	}

	γ := struct {
		OperationUserRcResponse `xml:"UserRcResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserRc", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserRcResponse, nil
}

// UserRedirect was auto-generated from WSDL.
func (p *pbxPortType) UserRedirect(UserRedirect *UserRedirect) (*UserRedirectResponse, error) {
	α := struct {
		OperationUserRedirectRequest `xml:"inno:UserRedirect"`
	}{
		OperationUserRedirectRequest{
			UserRedirect,
		},
	}

	γ := struct {
		OperationUserRedirectResponse `xml:"UserRedirectResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserRedirect", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserRedirectResponse, nil
}

// UserReroute was auto-generated from WSDL.
func (p *pbxPortType) UserReroute(UserReroute *UserReroute) (*UserRerouteResponse, error) {
	α := struct {
		OperationUserRerouteRequest `xml:"inno:UserReroute"`
	}{
		OperationUserRerouteRequest{
			UserReroute,
		},
	}

	γ := struct {
		OperationUserRerouteResponse `xml:"UserRerouteResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserReroute", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserRerouteResponse, nil
}

// UserRetrieve was auto-generated from WSDL.
func (p *pbxPortType) UserRetrieve(UserRetrieve *UserRetrieve) (*UserRetrieveResponse, error) {
	α := struct {
		OperationUserRetrieveRequest `xml:"inno:UserRetrieve"`
	}{
		OperationUserRetrieveRequest{
			UserRetrieve,
		},
	}

	γ := struct {
		OperationUserRetrieveResponse `xml:"UserRetrieveResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserRetrieve", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserRetrieveResponse, nil
}

// UserTransfer was auto-generated from WSDL.
func (p *pbxPortType) UserTransfer(UserTransfer *UserTransfer) (*UserTransferResponse, error) {
	α := struct {
		OperationUserTransferRequest `xml:"inno:UserTransfer"`
	}{
		OperationUserTransferRequest{
			UserTransfer,
		},
	}

	γ := struct {
		OperationUserTransferResponse `xml:"UserTransferResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserTransfer", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserTransferResponse, nil
}

// UserUUI was auto-generated from WSDL.
func (p *pbxPortType) UserUUI(UserUUI *UserUUI) (*UserUUIResponse, error) {
	α := struct {
		OperationUserUUIRequest `xml:"inno:UserUUI"`
	}{
		OperationUserUUIRequest{
			UserUUI,
		},
	}

	γ := struct {
		OperationUserUUIResponse `xml:"UserUUIResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/UserUUI", α, &γ); err != nil {
		return nil, err
	}
	return γ.UserUUIResponse, nil
}

// Version was auto-generated from WSDL.
func (p *pbxPortType) Version(Version *Version) (*VersionResponse, error) {
	α := struct {
		OperationVersionRequest `xml:"inno:Version"`
	}{
		OperationVersionRequest{
			Version,
		},
	}

	γ := struct {
		OperationVersionResponse `xml:"VersionResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("http://innovaphone.com/pbx/Version", α, &γ); err != nil {
		return nil, err
	}
	return γ.VersionResponse, nil
}

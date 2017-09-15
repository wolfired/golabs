package zerotier

type Resp struct {
	Members []Member `json:"members"`
}

type Member struct {
	Id                  string    `json:"id"`
	Type                string    `json:"type"`
	Clock               uint64    `json:"clock"`
	NetworkId           string    `json:"networkId"`
	NodeId              string    `json:"nodeId"`
	ControllerId        string    `json:"controllerId"`
	Hidden              string    `json:"hidden"`
	name                string    `json:"name"`
	online              bool      `json:"online"`
	description         string    `json:"description"`
	Config              Config    `json:"config"`
	LastOnline          uint64    `json:"lastOnline"`
	LastOffline         uint64    `json:"lastOffline"`
	PhysicalAddress     string    `json:"physicalAddress"`
	PhysicalLocation    []float32 `json:"physicalLocation"`
	ClientVersion       string    `json:"clientVersion"`
	OfflineNotifyDelay  uint64    `json:"offlineNotifyDelay"`
	ProtocolVersion     uint64    `json:"protocolVersion"`
	SupportsRulesEngine uint64    `json:"supportsRulesEngine"`
}

type Config struct {
	Id                           string   `json:"id"`
	Nwid                         string   `json:"nwid"`
	CreationTime                 uint64   `json:"creationTime"`
	Objtype                      string   `json:"objtype"`
	Revision                     uint64   `json:"revision"`
	Address                      string   `json:"address"`
	Authorized                   bool     `json:"authorized"`
	ActiveBridge                 bool     `json:"activeBridge"`
	Capabilities                 []string `json:"capabilities"`
	Identity                     string   `json:"identity"`
	RemoteTraceTarget            string   `json:"remoteTraceTarget"`
	IpAssignments                []string `json:"ipAssignments"`
	NoAutoAssignIps              bool     `json:"noAutoAssignIps"`
	Tags                         []string `json:"tags"`
	LastAuthorizedTime           uint64   `json:"lastAuthorizedTime"`
	LastAuthorizedCredentialType uint64   `json:"lastAuthorizedCredentialType"`
	LastAuthorizedCredential     uint64   `json:"lastAuthorizedCredential"`
	LastDeauthorizedTime         uint64   `json:"lastDeauthorizedTime"`
	PhysicalAddr                 string   `json:"physicalAddr"`
	VMajor                       uint64   `json:"vMajor"`
	VMinor                       uint64   `json:"vMinor"`
	VRev                         uint64   `json:"vRev"`
	VProto                       uint64   `json:"vProto"`
}

func (t *TierClient) NetworkMember() {

}

package zerotier

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Resp 响应
type Resp struct {
	Members []Member `json:"members"`
}

// Member 成员数据结构
type Member struct {
	ID                  string    `json:"id"`
	Type                string    `json:"type"`
	Clock               uint64    `json:"clock"`
	NetworkID           string    `json:"networkId"`
	NodeID              string    `json:"nodeId"`
	ControllerID        string    `json:"controllerId"`
	Hidden              string    `json:"hidden"`
	Name                string    `json:"name"`
	Online              bool      `json:"online"`
	Description         string    `json:"description"`
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

// Config 配置
type Config struct {
	ID                           string   `json:"id"`
	NWID                         string   `json:"nwid"`
	CreationTime                 uint64   `json:"creationTime"`
	Objtype                      string   `json:"objtype"`
	Revision                     uint64   `json:"revision"`
	Address                      string   `json:"address"`
	Authorized                   bool     `json:"authorized"`
	ActiveBridge                 bool     `json:"activeBridge"`
	Capabilities                 []string `json:"capabilities"`
	Identity                     string   `json:"identity"`
	RemoteTraceTarget            string   `json:"remoteTraceTarget"`
	IPAssignments                []string `json:"ipAssignments"`
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

// NetworkMember 查询网络成员
func (t *TierClient) NetworkMember() {
	req, _ := http.NewRequest("GET", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", nil)
	req.Header.Add("Authorization", "bearer xxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	resp, _ := http.DefaultClient.Do(req)

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	str := `{"members": ` + string(bytes) + `}`

	result := new(Resp)
	json.Unmarshal([]byte(str), result)
}

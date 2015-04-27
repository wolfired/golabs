package namesilo

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type ListDomainsResp struct {
	Request
	ListDomainsReply
}

type ListDomainsReply struct {
	Reply
	Domains []string `xml:"reply>domains>domain"`
}

func (self *SiloClient) ListDomains() *ListDomainsResp {
	resp, _ := http.Get(self.Url("listDomains", map[string]string{}))

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	result := new(ListDomainsResp)
	xml.Unmarshal(bytes, result)

	return result
}

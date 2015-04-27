package namesilo

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type DNSListRecordsResp struct {
	Request
	DnsListRecordsReply
}

type DnsListRecordsReply struct {
	Reply
	ResourceRecords []ResourceRecord `xml:"reply>resource_record"`
}

type ResourceRecord struct {
	RecordID string `xml:"record_id"`
	Type     string `xml:"type"`
	Host     string `xml:"host"`
	Value    string `xml:"value"`
	TTL      uint   `xml:"ttl"`
	Distance uint   `xml:"distance"`
}

func (self *SiloClient) DNSListRecords(domain string) *DNSListRecordsResp {
	resp, _ := http.Get(self.Url("dnsListRecords", map[string]string{"domain": domain}))

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	result := new(DNSListRecordsResp)
	xml.Unmarshal(bytes, result)

	return result
}

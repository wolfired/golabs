package namesilo

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DNSUpdateRecordResp struct {
	Request
	DNSUpdateRecordReply
}

type DNSUpdateRecordReply struct {
	Reply
	RecordID string `xml:"reply>record_id"`
}

func (self *SiloClient) DNSUpdateRecord(domain string, rrid string, rrhost string, rrvalue string, rrdistance uint, rrttl uint) *DNSUpdateRecordResp {
	resp, _ := http.Get(self.Url("dnsUpdateRecord", map[string]string{"domain": domain, "rrid": rrid, "rrhost": rrhost, "rrvalue": rrvalue, "rrdistance": fmt.Sprintf("%d", rrdistance), "rrttl": fmt.Sprintf("%d", rrttl)}))

	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	result := new(DNSUpdateRecordResp)
	xml.Unmarshal(bytes, result)

	return result
}

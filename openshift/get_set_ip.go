/*记录IP*/
package openshift

import (
	"fmt"
	"github.com/wolfired/golabs/namesilo"
	"net/http"
	"time"
)

type record struct {
	ip string    //要记录的IP
	ts time.Time //记录IP时间
}

var (
	Passwd  string //记录IP必须提供的密码，默认为空无法记录IP
	SiloCli *namesilo.SiloClient
)

var (
	record_map map[string]*record = make(map[string]*record, 8)
)

/*
记录IP到指定KEY，URL参数：passwd=密码，key=保存键值，domain=要同步更新A记录的域名（可选），host=要同步更新A记录的别名（可选）
*/
func SetIP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if nil != err {
		fmt.Fprintf(res, err.Error())
		return
	}

	passwd := req.Form.Get("passwd")
	key := req.Form.Get("key")

	if "" == Passwd || passwd != Passwd {
		fmt.Fprintf(res, "Wrong passwd for set ip")
		return
	}

	r := record_map[key]

	if nil == r {
		r = new(record)
		record_map[key] = r
	}

	r.ip = req.Header.Get("X-Forwarded-For") //内部端口跳转
	if "" == r.ip {
		r.ip = req.Host
	}
	r.ts = time.Now()

	fmt.Fprintf(res, "Set ip %s successfully", r.ip)

	domain := req.Form.Get("domain")
	host := req.Form.Get("host")
	if "" != domain && "" != host {
		updateARecord(key, domain, host)
	}
}

/*显示指定KEY下的IP下次更新的时间，URL参数：key=保存键值*/
func ShowNextTimestamp(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if nil != err {
		fmt.Fprintf(res, err.Error())
		return
	}

	key := req.Form.Get("key")

	r := record_map[key]

	if nil == r {
		fmt.Fprintf(res, "do not have a ip for the key %s", key)
		return
	}

	fmt.Fprintf(res, "%s", r.ts.UTC().In(time.FixedZone("+8", 28800)).Add(5*time.Minute).String())
}

/*获取指定KEY下的IP，URL参数：key=保存键值*/
func GetIP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if nil != err {
		fmt.Fprintf(res, err.Error())
		return
	}

	key := req.Form.Get("key")

	r := record_map[key]

	if nil == r {
		fmt.Fprintf(res, "do not have a ip for the key %s", key)
		return
	}

	fmt.Fprintf(res, "%s", r.ip)
}

//更新A记录
func updateARecord(key string, domain string, host string) {
	rr := SiloCli.DNSListRecords(domain).GetResourceRecordByHost(host + "." + domain)
	if nil == rr {
		return
	}

	r := record_map[key]
	if nil == r {
		return
	}

	if rr.Value == r.ip {
		return
	}

	SiloCli.DNSUpdateRecord(domain, rr.RecordID, host, r.ip, rr.Distance, rr.TTL)
}

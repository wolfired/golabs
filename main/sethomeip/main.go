package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/wolfired/golabs/namesilo"
)

var (
	/*SiloCli namesilo客户端*/
	SiloCli *namesilo.SiloClient
)

func init() {
	namesiloVersion, _ := strconv.Atoi(os.Getenv("NAMESILO_VERSION"))
	SiloCli = &namesilo.SiloClient{Version: uint(namesiloVersion), Type: os.Getenv("NAMESILO_TYPE"), Key: os.Getenv("NAMESILO_KEY")}
}

func main() {
	help := flag.Bool("help", false, "帮助")
	ip := flag.String("ip", "", "公网IP")
	domain := flag.String("domain", "wolfired.com", "域名")
	host := flag.String("host", "pi", "主机")
	flag.Parse()

	if *help || *ip == "" || *domain == "" || *host == "" {
		flag.Usage()
		return
	}

	rr := SiloCli.DNSListRecords(*domain).GetResourceRecordByHost(*host + "." + *domain)
	if nil == rr {
		return
	}

	if rr.Value == *ip {
		return
	}

	SiloCli.DNSUpdateRecord(*domain, rr.RecordID, *host, *ip, rr.Distance, rr.TTL)
}

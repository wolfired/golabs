package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	help := flag.Bool("help", false, "帮助")
	passwd := flag.String("passwd", "", "密码")
	key := flag.String("key", "", "标识键")
	domain := flag.String("domain", "wolfired.com", "域名")
	host := flag.String("host", "pi", "主机")
	flag.Parse()

	if *help || "" == *passwd || "" == *key || "" == *domain || "" == *host {
		flag.Usage()
		return
	}

	url := fmt.Sprintf("http://sloot.wolfired.com/set_ip?passwd=%s&key=%s&domain=%s&host=%s", *passwd, *key, *domain, *host)
	http.Get(url)
}

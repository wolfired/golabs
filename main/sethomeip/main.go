package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	passwd := os.Args[0]
	key := os.Args[1]
	domain := os.Args[2]
	host := os.Args[3]
	url := fmt.Sprintf("https://ddns-deriflow.rhcloud.com/set_ip?passwd=%s&key=%s&domain=%s&host=%s", passwd, key, domain, host)
	http.Get(url)
}

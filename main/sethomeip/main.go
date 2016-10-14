package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	passwd := os.Args[1]
	key := os.Args[2]
	domain := os.Args[3]
	host := os.Args[4]
	url := fmt.Sprintf("https://ddns-wolfired.rhcloud.com/set_ip?passwd=%s&key=%s&domain=%s&host=%s", passwd, key, domain, host)
	http.Get(url)
}

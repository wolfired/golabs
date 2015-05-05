package main

import (
	// "log"
	"net/http"
	"os"
	"time"
)

func main() {
	if 1 < len(os.Args) {
		setHomeIP()
		return
	}

	for true {
		setHomeIP()
		time.Sleep(5 * time.Minute)
	}
}

func setHomeIP() {
	_, err := http.Get("https://ddns-deriflow.rhcloud.com/set_home_ip?passwd=112358")
	if nil != err {
		// log.Fatalln(err)
	}

	/*	defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			log.Fatalln(err)
		}

		fmt.Println(string(body))*/
}

package main

import (
	// "log"
	"net/http"
	"os"
	"time"
)

func main() {
	for 1 == len(os.Args) {
		set_home_ip()
		time.Sleep(5 * time.Minute)
	}
}

func set_home_ip() {
	_, err := http.Get("https://ddns-deriflow.rhcloud.com//set_home_ip?passwd=112358")
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

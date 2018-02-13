package main

import (
	"flag"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
	"github.com/wolfired/golabs/idiotgs/session"
)

func main() {
	help := flag.Bool("help", false, "帮助")
	addr := flag.String("addr", "0.0.0.0:9999", "Server")
	flag.Parse()

	if *help || "" == *addr {
		flag.Usage()
		return
	}

	upgrader := ws.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	http.HandleFunc("/idiot", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if nil != err {
			log.Print("upgrade:", err)
			return
		}

		go session.MakeSession(c).Run()
	})

	http.ListenAndServe(*addr, nil)
}

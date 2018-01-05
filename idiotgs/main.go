package main

import (
	"flag"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
	"github.com/wolfired/golabs/idiotGS/session"
)

func main() {
	addr := flag.String("addr", "localhost:9999", "Server")

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

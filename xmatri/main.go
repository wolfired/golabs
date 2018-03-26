package main

import (
	"flag"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
	"github.com/wolfired/golabs/idiotgs/session"
)

const html string = `<html>

<head>
    <title>Hello Xmatri</title>
</head>

<body style="margin: 0px;">
	<script>
		let ws = new WebSocket("ws://localhost:9999/idiot");
		ws.addEventListener("open", () => {
			console.log("open");
		});
		let can = document.createElement("canvas");
		can.style = "width: 100px; height: 100px;";
		document.body.append(can);

		can.addEventListener("click", (e) => {
			console.log(e.x, e.y);
		});
		can.addEventListener("keydown",
			(e) => {
				console.log(e.keyCode);
			});
		can.addEventListener("keypress", (e) => {
			console.log(e.keyCode);
		});
	</script>
</body>

</html>`

func main() {
	help := flag.Bool("help", false, "帮助")
	addr := flag.String("addr", "0.0.0.0:9999", "Server")
	flag.Parse()

	if *help || "" == *addr {
		flag.Usage()
		return
	}

	upgrader := ws.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte(html))
	})

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


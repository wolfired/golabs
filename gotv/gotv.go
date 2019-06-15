package gotv

/*
cd gotv & tsc -m none -t ESNEXT gotv.ts & uglifyjs gotv.js -o gotv.js & move /Y gotv.js ..\testbed & cd ..\testbed & go run main.go & cd ..
*/
import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func serve() {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

	http.HandleFunc(flags.ws, func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if nil != err {
			log.Print("upgrade:", err)
			return
		}

		go newHub(conn).boot()
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(index.Bytes())
	})

	http.ListenAndServe(flags.addr+":"+flags.port, nil)
}

//Boot Boot
func Boot() {
	parse()

	html()

	serve()
}

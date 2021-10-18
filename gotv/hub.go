package gotv

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wolfired/golabs/gotv/frame"
)

type hub struct {
	conn *websocket.Conn
	ctrl chan byte
	buf  chan []byte
}

func newHub(conn *websocket.Conn) *hub {
	h := new(hub)
	h.conn = conn
	h.ctrl = make(chan byte)
	h.buf = make(chan []byte)
	return h
}

func (h *hub) recv() {
	for {
		mt, raw, err := h.conn.ReadMessage()
		if nil != err {
			log.Println("recv:", err, mt)
			close(h.ctrl)
			break
		}
		switch mt {
		case websocket.TextMessage:
			log.Printf("TextMessage: %s\n", raw)
			h.buf <- raw
		case websocket.BinaryMessage:
			log.Printf("BinaryMessage:")
			h.buf <- raw
		case websocket.CloseMessage:
			log.Printf("CloseMessage:")

		case websocket.PingMessage:
			log.Printf("PingMessage:")

		case websocket.PongMessage:
			log.Printf("PongMessage:")

		}
	}
}

func (h *hub) send() {
	for {
		select {
		case buf := <-h.buf:
			err := h.conn.WriteMessage(websocket.BinaryMessage, buf)
			if nil != err {
				log.Println("send:", err)
				close(h.ctrl)
				break
			}
		}
	}
}

func (h *hub) boot() {
	log.Println("hub boot")

	defer h.conn.Close()

	go h.recv()

	go h.send()

	go h.draw()

	<-h.ctrl

	log.Println("hub halt")
}

func (h *hub) draw() {
	t := time.Tick(time.Millisecond * 41)

	i := 0
	a := frame.FromFile(flags.spritesheet, flags.sprite_wid, flags.sprite_hei, flags.sprite_count)

	for {
		select {
		case <-t:
			h.buf <- []byte(a.FrameAt(i%a.Len(), flags.enable_zip))
			i++
			// h.buf <- []byte("hi")
		}
	}
}

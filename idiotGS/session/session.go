package session

import (
	"log"

	ws "github.com/gorilla/websocket"
)

// Session 会话
type Session struct {
	Sock *ws.Conn
	Buf  chan []byte
	Done chan struct{}
}

// MakeSession 创建会话
func MakeSession(c *ws.Conn) (s *Session) {
	s = new(Session)
	s.Sock = c
	s.Buf = make(chan []byte)
	s.Done = make(chan struct{})
	return
}

// Run 运行
func (s *Session) Run() {
	defer s.Sock.Close()

	go func() {
		for {
			mt, dat, err := s.Sock.ReadMessage()
			if nil != err {
				log.Println("read:", err)
				close(s.Done)
				break
			}
			switch mt {
			case ws.TextMessage:
				log.Println(dat)
			case ws.BinaryMessage:
				log.Println(string(dat))
			}
		}
	}()

	go func() {
		for {
			select {
			case buf := <-s.Buf:
				err := s.Sock.WriteMessage(ws.BinaryMessage, buf)
				if nil != err {
					log.Println("write:", err)
					close(s.Done)
					break
				}
			}
		}
	}()

	<-s.Done

	log.Println("session closed")
}

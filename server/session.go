package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/wolfired/golabs/server/logic"
)

type Session struct {
	conn net.Conn
}

func (s *Session) Setup(c net.Conn) *Session {
	s.conn = c
	return s
}

func (s *Session) Run() {
	ra := s.conn.RemoteAddr().String()

	bfio := bufio.NewReaderSize(s.conn, 1024)

	for {
		line, err := bfio.ReadString(0)
		if nil == err {
			logic.Saver[0]()
			fmt.Println(ra, "->", line)
		} else {
			break
		}
	}
}

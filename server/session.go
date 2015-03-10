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

func (self *Session) Setup(c net.Conn) *Session {
	self.conn = c
	return self
}

func (self *Session) Run() {
	ra := self.conn.RemoteAddr().String()

	bfio := bufio.NewReaderSize(self.conn, 1024)

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

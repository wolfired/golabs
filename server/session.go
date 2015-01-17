package server

import (
	"bufio"
	"fmt"
	"net"
)

type Session struct {
	conn net.Conn
}

func (this *Session) Setup(c net.Conn) *Session {
	this.conn = c
	return this
}

func (this *Session) Run() {
	ra := this.conn.RemoteAddr().String()

	bfio := bufio.NewReaderSize(this.conn, 1024)

	for {
		line, err := bfio.ReadString(0)
		if nil == err {
			fmt.Println(ra, "->", line)
		} else {
			break
		}
	}
}

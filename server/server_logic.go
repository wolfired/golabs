package server

import (
	"bufio"
	"fmt"
	"net"
)

type LogicServer struct {
	Server
}

func (l *LogicServer) Run() {
	ln, err := net.Listen(l.Net, l.Addr)
	if nil != err {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if nil != err {
			fmt.Println(err)
			continue
		}
		ra := conn.RemoteAddr().String()
		fmt.Println(ra)

		bfio := bufio.NewReaderSize(conn, 1024)
		for {
			line, err := bfio.ReadString(0)
			if nil == err {
				fmt.Println(ra, "->", line)
			} else {
				break
			}
		}
	}
}
package server

import (
	"net"
	"bufio"
	"fmt"
)

type LogicServer struct {
	Net string
	Addr string
}

func (this *LogicServer) Run() {
	ln, err := net.Listen(this.Net, this.Addr)
	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		bfio := bufio.NewReaderSize(conn, 1024)
		for {
			line, err := bfio.ReadString(0)
			if nil == err {
				fmt.Println("->", line)
				} else {
					break
				}
			}
	}
}

package server

import (
	"fmt"
	"net"
)

type GateServer struct {
	Server
}

func (self *GateServer) Run() {
	ln, err := net.Listen(self.Net, self.Addr)
	if nil != err {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if nil != err {
			fmt.Println(err)
			continue
		}

		go new(Session).Setup(conn).Run()
	}
}

package server

import (
	"fmt"
	"net"
)

type GateServer struct {
	Server
}

func (g *GateServer) Run() {
	ln, err := net.Listen(g.Net, g.Addr)

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

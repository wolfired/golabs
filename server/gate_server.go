package server

import (
	"net"
)

type GateServer struct {
	Net  string
	Addr string
}

func (this *GateServer) Run() {
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

		go new(Session).Setup(conn).Run()
	}
}

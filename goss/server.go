package goss

import (
	"fmt"
	"net"
)

func serve() {
	listen, err := net.Listen(flags.network, flags.host+":"+flags.port)

	if nil != err {
		fmt.Println(err)
	}

	for {
		conn, err := listen.Accept()

		if nil != err {
			fmt.Println(err)
			continue
		}

		go newHub(conn).boot()
	}
}

//BootServer BootServer
func BootServer() {
	parse()
	serve()
}

package goss

import (
	"fmt"
	"net"
)

func serve() {
	listen, err := net.Listen(flags.network, flags.address)

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

//Boot Boot
func Boot() {
	parse()

	go pfr()

	serve()
}

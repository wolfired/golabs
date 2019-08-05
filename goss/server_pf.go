package goss

import (
	"fmt"
	"net"
)

const (
	key = `<policy-file-request/>`
	pf  = `<?xml version="1.0"?><!DOCTYPE cross-domain-policy SYSTEM 'http://www.macromedia.com/xml/dtds/cross-domain-policy.dtd'><cross-domain-policy><allow-access-from domain="*" to-ports="*" /></cross-domain-policy>`
)

func servePF() {
	listen, err := net.Listen(flags.network, flags.host+":"+flags.pfPort)

	if nil != err {
		fmt.Println(err)
	}

	for {
		conn, err := listen.Accept()

		if nil != err {
			fmt.Println(err)
			continue
		}

		bs := make([]byte, len(key))
		conn.Read(bs)

		if key == string(bs) {
			conn.Write([]byte(pf))
			conn.Write([]byte{0})
		}
	}
}

//BootServerPF BootServerPF
func BootServerPF() {
	parse()
	go servePF()
}

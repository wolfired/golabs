package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Println(conn.RemoteAddr().String())
	conn.Close()
}

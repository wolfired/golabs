package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	time.Tick(d)

	fmt.Println(conn.RemoteAddr().String())
	fmt.Printf("%T\n", conn)

	b := make([]byte, 0, 1204)
	conn.Read()
	fmt.Println(len(b))
}

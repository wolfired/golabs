package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	help := *flag.Bool("help", false, "Help")
	address := *flag.String("address", "0.0.0.0:3980", "Address")

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	listen, err := net.Listen("tcp", address)

	if nil != err {
		fmt.Println(err)
	}

	fmt.Printf("%s waitting for connect @ %s\n", time.Now().Format("2006-01-02 15:04:05"), address)
	id := -1

	for {
		conn, err := listen.Accept()

		if nil != err {
			fmt.Println(err)
			continue
		}

		id++

		go func(id int, conn net.Conn) {
			fmt.Printf("%s conn %08d connected\n", time.Now().Format("2006-01-02 15:04:05"), id)

			chanCmd := make(chan byte)
			chanBuf := make(chan []byte, 64)

			// recv
			go func() {
				for {
					buf := make([]byte, 64)
					n, err := conn.Read(buf)
					if nil != err {
						fmt.Printf("%s conn %08d R error: %s\n", time.Now().Format("2006-01-02 15:04:05"), id, err.Error())
						chanCmd <- 0
						break
					}
					chanBuf <- buf[:n]
				}
			}()

			// send
			go func() {
				for {
					select {
					case buf := <-chanBuf:
						{
							fmt.Printf("%s conn %08d recv and send: %d bytes\n", time.Now().Format("2006-01-02 15:04:05"), id, len(buf))
							_, err = conn.Write(buf)
							if nil != err {
								fmt.Printf("%s conn %08d W error: %s\n", time.Now().Format("2006-01-02 15:04:05"), id, err.Error())
								chanCmd <- 0
								break
							}
						}
					}
				}
			}()

			// ctrl
			for {
				select {
				case cmd := <-chanCmd:
					{
						switch cmd {
						case 0:
							{
								fmt.Printf("%s conn %08d disconnected\n", time.Now().Format("2006-01-02 15:04:05"), id)
								conn.Close()
								break
							}
						}
					}
				}
			}
		}(id, conn)
	}
}

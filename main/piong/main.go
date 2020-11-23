package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	help := flag.Bool("help", false, "Help")
	mode := flag.String("mode", "client", "Mode: server or client")
	address := flag.String("address", "0.0.0.0:3980", "Address")
	bsize := flag.Int("bsize", 4096, "Buf size")

	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if "server" == *mode {
		server(*address, *bsize)
	} else if "client" == *mode {
		client(*address, *bsize)
	} else {
		flag.Usage()
	}
}

func server(address string, bsize int) {
	listen, err := net.Listen("tcp", address)

	if nil != err {
		fmt.Println(err)
		return
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
			defer conn.Close()

			fmt.Printf("%s conn %08d connected\n", time.Now().Format("2006-01-02 15:04:05"), id)

			chanCmd := make(chan byte)
			chanBuf := make(chan []byte, bsize)

			// recv
			go func() {
				for {
					buf := make([]byte, bsize)
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

func client(address string, bsize int) {
	conn, err := net.Dial("tcp", address)

	if nil != err {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	chanCmd := make(chan byte)
	chanBuf := make(chan []byte, bsize)

	// recv
	go func() {
		for {
			buf := make([]byte, bsize)
			n, err := conn.Read(buf)
			if nil != err {
				fmt.Printf("%s client R error: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
				chanCmd <- 0
				break
			}
			fmt.Printf("%s client recv: %s in %d bytes\n", time.Now().Format("2006-01-02 15:04:05"), string(buf[:n]), n)
			fmt.Printf("%s please input: ", time.Now().Format("2006-01-02 15:04:05"))
		}
	}()

	// send
	go func() {
		for {
			select {
			case buf := <-chanBuf:
				{
					fmt.Printf("\r%s client send: %s in %d bytes\n", time.Now().Format("2006-01-02 15:04:05"), string(buf), len(buf))
					_, err = conn.Write(buf)
					if nil != err {
						fmt.Printf("%s client W error: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
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
						fmt.Printf("%s client disconnected\n", time.Now().Format("2006-01-02 15:04:05"))
						conn.Close()
						break
					}
				}
			}
		default:
			{
				fmt.Printf("%s please input: ", time.Now().Format("2006-01-02 15:04:05"))

				buf := make([]byte, bsize)
				n, err := os.Stdin.Read(buf)

				if nil != err {
					fmt.Printf("%s input error: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
					continue
				}

				if 0 < len(buf[:n-1]) {
					chanBuf <- buf[:n-1]
				}
			}
		}
	}

}

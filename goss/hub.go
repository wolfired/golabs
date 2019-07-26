package goss

import (
	"log"
	"net"

	"github.com/wolfired/as2go/flash/utils"
)

type hub struct {
	conn net.Conn
	cmd  chan byte
	buf  chan []byte
}

func newHub(conn net.Conn) *hub {
	h := new(hub)
	h.conn = conn
	h.cmd = make(chan byte)
	h.buf = make(chan []byte)
	return h
}

func (h *hub) ctrl() {
	for {
		select {
		case cmd := <-h.cmd:
			{
				switch cmd {
				case 0:
					{
						h.conn.Close()
						return
					}
				}
			}
		}
	}
}

func (h *hub) send() {
	for {
		select {
		case buf := <-h.buf:
			{
				h.conn.Write(buf)
			}
		}
	}
}

func (h *hub) recv() {
	ori := make([]byte, 64)
	ba := utils.NewByteArray(ori)
	buf := ori[:2]

	for {
		need := len(buf)
		n, err := h.conn.Read(buf)

		if nil != err {
			h.cmd <- 0
			return
		}

		if n != need {
			buf = buf[n:need]
			continue
		}

		ba.SetPosition(0)
		ba.SetLength(uint(need))

		if 2 == need {
			size, _ := ba.ReadUnsignedShort()
			buf = buf[:size]
			continue
		}

		s, _ := ba.ReadUTF()
		log.Println(s)

		buf = ori[0:2]
	}
}

func (h *hub) boot() {
	log.Println("hub boot")

	go h.send()
	go h.recv()
	h.ctrl()

	log.Println("hub halt")
}

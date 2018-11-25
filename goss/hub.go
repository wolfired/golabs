package goss

import (
	"encoding/binary"
	"log"
	"net"
)

type hub struct {
	conn net.Conn
	ctrl chan byte
	buf  chan []byte
	cmd  chan []byte
}

func newHub(conn net.Conn) *hub {
	h := new(hub)
	h.conn = conn
	h.ctrl = make(chan byte)
	h.buf = make(chan []byte)
	h.cmd = make(chan []byte)
	return h
}

func (h *hub) recv() {
	for {
		size := make([]byte, 2)

		if nil != h.read(size) {
			break
		}

		if 0 == size[0] && 0 == size[1] {
			continue
		}

		cmd := make([]byte, binary.BigEndian.Uint16(size))

		if nil != h.read(cmd) {
			break
		}

		h.cmd <- cmd
	}
}

func (h *hub) read(buf []byte) error {
	n := 0
	l := len(buf)

	for 0 < l {
		rn, err := h.conn.Read(buf[n:])

		if nil != err {
			log.Println("recv:", err)
			close(h.ctrl)
			return err
		}

		n += rn
		l -= rn
	}

	return nil
}

func (h *hub) send() {
	for {
		select {
		case buf := <-h.buf:
			_, err := h.conn.Write(buf)

			if nil != err {
				log.Println("send:", err)
				close(h.ctrl)
				break
			}
		}
	}
}

func (h *hub) handle() {
	for {
		select {
		case cmd := <-h.cmd:
			bs := make([]byte, 2)
			binary.BigEndian.PutUint16(bs, uint16(len(cmd)))
			h.buf <- bs
			h.buf <- cmd
		}
	}
}

func (h *hub) boot() {
	log.Println("hub boot")

	defer h.conn.Close()

	go h.send()

	go h.handle()

	go h.recv()

	<-h.ctrl

	log.Println("hub halt")
}

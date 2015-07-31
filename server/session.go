package server

import (
	"bufio"
	"fmt"
	"net"
)

type Session struct {
	conn net.Conn
}

func (s *Session) Setup(c net.Conn) *Session {
	s.conn = c
	return s
}

func (s *Session) Run() {

}

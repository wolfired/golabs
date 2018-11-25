package gotv

import (
	"bytes"
	"encoding/binary"
)

const (
	ctrlHalt = 0xFE
)

type codec interface {
	encode(buf *bytes.Buffer)
	decode(buf *bytes.Buffer)
}

type msg struct {
	head uint32
}

func newMsg(ctrl byte) *msg {
	m := new(msg)
	m.head = uint32(ctrl)
	return m
}

func (m *msg) encode(buf *bytes.Buffer) {

}

func (m *msg) decode(buf *bytes.Buffer) {
	m.head = binary.BigEndian.Uint32(buf.Next(4))
}

func (m *msg) ctrl() byte {
	return byte(m.head & 0xFF)
}

func (m *msg) size() uint32 {
	return m.head >> 0x8
}

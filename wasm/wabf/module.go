package wabf

import (
	"fmt"
)

const (
	magic = `\0asm`
)

//Decode TODO
func Decode(raw []byte) interface{} {
	m := module{}
	m.decode(raw)
	// fmt.Println(m)
	fmt.Println(m.sections[5])
	// fmt.Println(m.sections[0].payloadType)
	return m
}

type decoder interface {
	decode(raw []byte) uint
}

type module struct {
	magic    string
	version  uint32
	sections []section
}

func (m *module) decode(raw []byte) uint {
	m.magic = string(raw[:4])
	raw = raw[4:]

	m.version = uint32(raw[3])<<24 | uint32(raw[2])<<16 | uint32(raw[1])<<8 | uint32(raw[0])
	raw = raw[4:]

	count := calcSectionCount(raw)
	m.sections = make([]section, count)

	for i := uint(0); i < count; i++ {
		used := m.sections[i].decode(raw)
		raw = raw[used:]
	}

	return 0
}

func calcSectionCount(raw []byte) uint {
	payloadLen, count := varuint32(0), uint(0)

	for i := uint(0); i < uint(len(raw)); i += uint(payloadLen) {
		payloadLen += varuint32(payloadLen.decode(raw[i+1:]) + 1)
		count++
	}

	return count
}

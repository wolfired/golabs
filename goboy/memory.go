package goboy

type memory [1*1024*1024]byte

func (m *memory)writeByte(addr uint, val byte)  {
	m[addr] = val
}

func (m *memory)readByte(addr uint) byte {
	return m[addr]
}
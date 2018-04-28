package goboy

type memory [1 * 1024 * 1024]byte

func (m *memory) setByte(addr uint, val byte) {
	m[addr] = val
}

func (m *memory) getByte(addr uint) byte {
	return m[addr]
}

package goboy

const (
	flagZ byte = 0x7
	flagN byte = 0x6
	flagH byte = 0x5
	flagC byte = 0x4
)

type register struct{
	a, f byte
	b, c byte
	d, e byte
	h, l byte
	sp, pc uint16
}
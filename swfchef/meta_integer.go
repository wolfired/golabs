package swfchef

import ()

type ui8 uint8
type ui16 uint16
type ui32 uint32
type ui8n []uint8
type ui16n []uint16
type ui24n []uint32
type ui32n []uint32
type ui64n []uint64

type si8 int8
type si16 int16
type si32 int32
type si8n []int8
type si16n []int16

func b2ui8(b []byte) (r ui8) {
	r = ui8(b[0])
	return
}

func b2ui16(b []byte) (r ui16) {
	r = ui16(b[0]) | ui16(b[1])<<8
	return
}

func b2ui32(b []byte) (r ui32) {
	r = ui32(b[0]) | ui32(b[1])<<8 | ui32(b[2])<<16 | ui32(b[3])<<24
	return
}

func b2ui8n(b []byte, n uint16) (r []ui8) {
	r = make([]ui8, n)
	for i := uint16(0); i < n; i += 1 {
		r[i] = b2ui8(b)
	}
	return
}

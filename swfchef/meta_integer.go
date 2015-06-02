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

func raw2ui8(raw []byte) (r ui8) {
	r = ui8(raw[0])
	return
}

func raw2ui16(raw []byte) (r ui16) {
	r = ui16(raw[0]) | ui16(raw[1])<<8
	return
}

func raw2ui32(raw []byte) (r ui32) {
	r = ui32(raw[0]) | ui32(raw[1])<<8 | ui32(raw[2])<<16 | ui32(raw[3])<<24
	return
}

func raw2ui8n(raw []byte, n uint16) (r []ui8) {
	r = make([]ui8, n)
	for i := uint16(0); i < n; i += 1 {
		r[i] = raw2ui8(raw)
	}
	return
}

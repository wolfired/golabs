package swfchef

import ()

type ui8 uint8
type ui16 uint16
type ui24 uint32
type ui32 uint32
type ui64 uint64
type ui8n []ui8
type ui16n []ui16
type ui24n []ui24
type ui32n []ui32
type ui64n []ui64

type si8 int8
type si16 int16
type si32 int32
type si8n []si8
type si16n []si16

func raw2ui8(raw []byte) (r ui8) {
	r = ui8(raw[0])
	return
}

func raw2ui16(raw []byte) (r ui16) {
	r = ui16(raw[1])<<8 | ui16(raw[0])
	return
}

func raw2ui24(raw []byte) (r ui24) {
	r = ui24(raw[2])<<16 | ui24(raw[1])<<8 | ui24(raw[0])
	return
}

func raw2ui32(raw []byte) (r ui32) {
	r = ui32(raw[3])<<24 | ui32(raw[2])<<16 | ui32(raw[1])<<8 | ui32(raw[0])
	return
}

func raw2ui64(raw []byte) (r ui64) {
	r = ui64(raw[7])<<56 | ui64(raw[6])<<48 | ui64(raw[5])<<40 | ui64(raw[4])<<32 | ui64(raw[3])<<24 | ui64(raw[2])<<16 | ui64(raw[1])<<8 | ui64(raw[0])
	return
}

func raw2ui8n(raw []byte, n ui32) (r ui8n) {
	r = make(ui8n, n)
	for i := ui32(0); i < n; i += 1 {
		r[i] = raw2ui8(raw[i:])
	}
	return
}

func raw2ui16n(raw []byte, n ui32) (r ui16n) {
	r = make(ui16n, n)
	for i := ui32(0); i < n; i += 1 {
		r[i] = raw2ui16(raw[2*i:])
	}
	return
}

func raw2ui24n(raw []byte, n ui32) (r ui24n) {
	r = make(ui24n, n)
	for i := ui32(0); i < n; i += 1 {
		r[i] = raw2ui24(raw[3*i:])
	}
	return
}

func raw2ui32n(raw []byte, n ui32) (r ui32n) {
	r = make(ui32n, n)
	for i := ui32(0); i < n; i += 1 {
		r[i] = raw2ui32(raw[4*i:])
	}
	return
}

func raw2ui64n(raw []byte, n ui32) (r ui64n) {
	r = make(ui64n, n)
	for i := ui32(0); i < n; i += 1 {
		r[i] = raw2ui64(raw[8*i:])
	}
	return
}

func raw2si8(raw []byte) (r si8) {
	r = si8(raw[0])
	return
}

func raw2si16(raw []byte) (r si16) {
	r = si16(raw[1])<<8 | si16(ui16(raw[0]))
	return
}

func raw2si32(raw []byte) (r si32) {
	r = si32(raw[3])<<24 | si32(ui32(raw[2])<<16|ui32(raw[1])<<8|ui32(raw[0]))
	return
}

func raw2si8n(raw []byte, n ui32) (r si8n) {
	r = make(si8n, n)
	for i := ui32(0); i < n; i += 1 {
		r[i] = raw2si8(raw[i:])
	}
	return
}

func raw2si16n(raw []byte, n ui32) (r si16n) {
	r = make(si16n, n)
	for i := ui32(0); i < n; i += 1 {
		r[i] = raw2si16(raw[2*i:])
	}
	return
}

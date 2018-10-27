package swfchef

import (
	// "fmt"
	"testing"
)

var (
	zero []byte = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	one  []byte = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	raw  []byte = []byte{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8, 0xF7, 0xF6, 0xF5, 0xF4, 0xF3, 0xF2, 0xF1, 0xF0}
)

func TestRaw2ui8(t *testing.T) {
	if 0x00 != raw2ui8(zero) {
		t.Fail()
	}

	if 0xFF != raw2ui8(one) {
		t.Fail()
	}
}

func TestRaw2ui16(t *testing.T) {
	if 0x00 != raw2ui16(zero) {
		t.Fail()
	}

	if 0xFFFF != raw2ui16(one) {
		t.Fail()
	}
}

func TestRaw2ui32(t *testing.T) {
	if 0x00 != raw2ui32(zero) {
		t.Fail()
	}

	if 0xFFFFFFFF != raw2ui32(one) {
		t.Fail()
	}
}

func TestRaw2ui64(t *testing.T) {
	if 0x00 != raw2ui64(zero) {
		t.Fail()
	}

	if 0xFFFFFFFFFFFFFFFF != raw2ui64(one) {
		t.Fail()
	}
}

func TestRaw2ui8n(t *testing.T) {
	lf := ui8n{0xFF, 0xFE}
	ri := raw2ui8n(raw, 2)

	if len(lf) != len(ri) {
		t.Fail()
	}
	for i, _ := range ri {
		if lf[i] != ri[i] {
			t.Fail()
		}
	}
}

func TestRaw2ui16n(t *testing.T) {
	lf := ui16n{0xFEFF, 0xFCFD}
	ri := raw2ui16n(raw, 2)

	if len(lf) != len(ri) {
		t.Fail()
	}
	for i, _ := range ri {
		if lf[i] != ri[i] {
			t.Fail()
		}
	}
}

func TestRaw2ui24n(t *testing.T) {
	lf := ui24n{0xFDFEFF, 0xFAFBFC}
	ri := raw2ui24n(raw, 2)

	if len(lf) != len(ri) {
		t.Fail()
	}
	for i, _ := range ri {
		if lf[i] != ri[i] {
			t.Fail()
		}
	}
}

func TestRaw2ui32n(t *testing.T) {
	lf := ui32n{0xFCFDFEFF, 0xF8F9FAFB}
	ri := raw2ui32n(raw, 2)

	if len(lf) != len(ri) {
		t.Fail()
	}
	for i, _ := range ri {
		if lf[i] != ri[i] {
			t.Fail()
		}
	}
}

func TestRaw2ui64n(t *testing.T) {
	lf := ui64n{0xF8F9FAFBFCFDFEFF, 0xF0F1F2F3F4F5F6F7}
	ri := raw2ui64n(raw, 2)

	if len(lf) != len(ri) {
		t.Fail()
	}
	for i, _ := range ri {
		if lf[i] != ri[i] {
			t.Fail()
		}
	}
}

func TestRaw2si8(t *testing.T) {
	if -128 != raw2si8([]byte{0x80}) {
		t.Fail()
	}

	if -1 != raw2si8([]byte{0xFF}) {
		t.Fail()
	}

	if 0 != raw2si8([]byte{0x00}) {
		t.Fail()
	}

	if 1 != raw2si8([]byte{0x01}) {
		t.Fail()
	}

	if 127 != raw2si8([]byte{0x7F}) {
		t.Fail()
	}
}

func TestRaw2si16(t *testing.T) {
	if -32768 != raw2si16([]byte{0x00, 0x80}) {
		t.Fail()
	}

	if -1 != raw2si16([]byte{0xFF, 0xFF}) {
		t.Fail()
	}

	if 0 != raw2si16([]byte{0x00, 0x00}) {
		t.Fail()
	}

	if 1 != raw2si16([]byte{0x01, 0x00}) {
		t.Fail()
	}

	if 32767 != raw2si16([]byte{0xFF, 0x7F}) {
		t.Fail()
	}
}

func TestRaw2si32(t *testing.T) {
	if -2147483648 != raw2si32([]byte{0x00, 0x00, 0x00, 0x80}) {
		t.Fail()
	}

	if -1 != raw2si32([]byte{0xFF, 0xFF, 0xFF, 0xFF}) {
		t.Fail()
	}

	if 0 != raw2si32([]byte{0x00, 0x00, 0x00, 0x00}) {
		t.Fail()
	}

	if 1 != raw2si32([]byte{0x01, 0x00, 0x00, 0x00}) {
		t.Fail()
	}

	if 2147483647 != raw2si32([]byte{0xFF, 0xFF, 0xFF, 0x7F}) {
		t.Fail()
	}
}

func TestRaw2si8n(t *testing.T) {
	lf := si8n{-1, -2}
	ri := raw2si8n(raw, 2)

	if len(lf) != len(ri) {
		t.Fail()
	}
	for i, _ := range ri {
		if lf[i] != ri[i] {
			t.Fail()
		}
	}
}

func TestRaw2si16n(t *testing.T) {
	lf := si16n{-1, -2}
	ri := raw2si16n([]byte{0xFF, 0xFF, 0xFE, 0xFF}, 2)

	if len(lf) != len(ri) {
		t.Fail()
	}
	for i, _ := range ri {
		if lf[i] != ri[i] {
			t.Fail()
		}
	}
}

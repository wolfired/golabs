package swfchef

import (
	"testing"
)

var (
	raw []byte = []byte{0xFF, 0xFE}
)

func TestRaw2ui8(t *testing.T) {
	if 0xFF != raw2ui8(raw) {
		t.Fail()
	}
}

func TestRaw2ui16(t *testing.T) {
	if 0xFEFF != raw2ui16(raw) {
		t.Fail()
	}
}

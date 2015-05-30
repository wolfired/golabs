package swfchef

import (
	"bytes"
)

func ui8(b *bytes.Buffer) (r uint8) {
	r, _ = b.ReadByte()
	return
}

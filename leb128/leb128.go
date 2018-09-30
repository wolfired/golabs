package leb128

//DecodeU TODO
func DecodeU(raw []byte) (uint, uint) {
	r := uint(0)
	i := uint(0)

	for {
		r |= (uint(raw[i]&0x7F) << uint(i*7))

		if 0 == (raw[i] & 0x80) {
			return r, i + 1
		}

		i++
	}
}

//DecodeS TODO
func DecodeS(raw []byte) (int, uint) {
	r := int(0)
	i := uint(0)

	for {
		r |= (int(raw[i]&0x7F) << uint(i*7))

		if 0 == (raw[i] & 0x80) {
			return r, i + 1
		}

		i++
	}
}

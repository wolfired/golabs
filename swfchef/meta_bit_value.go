package swfchef

type sbn int64
type ubn uint64
type fbn float64

var (
	bitMasks [8]byte = [8]byte{0xFF, 0x7F, 0x3F, 0x1F, 0x0F, 0x07, 0x03, 0x01}
)

func raw2ubn(raw []byte, offset uint8, n uint8) (r ubn) {
	head_byte_offset, head_bit_offset := offset/8, offset%8
	tail_byte_offset, tail_bit_offset := (offset+n)/8, (offset+n)%8

	if 8 >= head_bit_offset+n {
		r = ubn(uint64(raw[head_byte_offset]) << head_bit_offset >> (8 - n))
		return
	}

	c := uint64(raw[head_byte_offset]<<head_bit_offset) << (n - 8)

	for head_byte_offset += 1; head_byte_offset < tail_byte_offset; head_byte_offset += 1 {
		c |= uint64(raw[head_byte_offset]) << ((tail_byte_offset-head_byte_offset-1)*8 + tail_bit_offset)
	}

	c |= uint64(raw[tail_byte_offset] >> (8 - tail_bit_offset))

	r = ubn(c)

	return
}

func raw2sbn(raw []byte, offset uint8, n uint8) (r sbn) {
	head_byte_offset, head_bit_offset := offset/8, offset%8
	tail_byte_offset, tail_bit_offset := (offset+n)/8, (offset+n)%8

	if 8 >= head_bit_offset+n {
		r = sbn(int8(raw[head_byte_offset]) << head_bit_offset >> (8 - n))
		return
	}

	c := int64(raw[head_byte_offset]<<head_bit_offset) << (n - 8)

	for head_byte_offset += 1; head_byte_offset < tail_byte_offset; head_byte_offset += 1 {
		c |= int64(uint64(raw[head_byte_offset]) << ((tail_byte_offset-head_byte_offset-1)*8 + tail_bit_offset))
	}

	c |= int64(raw[tail_byte_offset] >> (8 - tail_bit_offset))

	r = sbn(c)

	return
}

func Raw2sbn(b []byte, offset uint8, n uint8) (r sbn) {
	return raw2sbn(b, offset, n)
}

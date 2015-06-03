package swfchef

type ubn ui32
type sbn si32
type fbn float64

func raw2ubn(raw []byte, offset ubn, n ubn) (r ubn) {
	head_byte_offset, head_bit_offset := offset/8, offset%8
	tail_byte_offset, tail_bit_offset := (offset+n)/8, (offset+n)%8

	if 8 >= head_bit_offset+n {
		r = ubn(raw[head_byte_offset]) << head_bit_offset >> (8 - n)
		return
	}

	c := ubn(raw[head_byte_offset]<<head_bit_offset) << (n - 8)

	for head_byte_offset += 1; head_byte_offset < tail_byte_offset; head_byte_offset += 1 {
		c |= ubn(raw[head_byte_offset]) << ((tail_byte_offset-head_byte_offset-1)*8 + tail_bit_offset)
	}

	c |= ubn(raw[tail_byte_offset] >> (8 - tail_bit_offset))

	r = ubn(c)

	return
}

func raw2sbn(raw []byte, offset ubn, n ubn) (r sbn) {
	head_byte_offset, head_bit_offset := offset/8, offset%8
	tail_byte_offset, tail_bit_offset := (offset+n)/8, (offset+n)%8

	if 8 >= head_bit_offset+n {
		r = sbn(raw[head_byte_offset]) << head_bit_offset >> (8 - n)
		return
	}

	c := sbn(raw[head_byte_offset]<<head_bit_offset) << (n - 8)

	for head_byte_offset += 1; head_byte_offset < tail_byte_offset; head_byte_offset += 1 {
		c |= sbn(raw[head_byte_offset]) << ((tail_byte_offset-head_byte_offset-1)*8 + tail_bit_offset)
	}

	c |= sbn(raw[tail_byte_offset] >> (8 - tail_bit_offset))

	r = sbn(c)

	return
}

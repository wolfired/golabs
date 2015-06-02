package swfchef

type fixed32 float32
type fixed16 float32

func raw2fixed16(raw []byte) (r fixed16) {
	r = fixed16(raw[1] + (raw[0]>>4)/16.0 + (raw[0]&0x0F)/16.0/16.0)
	return
}

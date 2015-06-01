package swfchef

type rectangle struct {
	xmin sbn
	xmax sbn
	ymin sbn
	ymax sbn
}

type fixed float32
type fixed8 float32

func b2rectangle(b []byte) (r rectangle) {
	r = rectangle{}
	n := uint8(b2sbn(b, 0, 5))
	r.xmin = b2sbn(b, 5, n)
	r.xmax = b2sbn(b, 5+n, n)
	r.ymin = b2sbn(b, 5+2*n, n)
	r.ymax = b2sbn(b, 5+2*n, n)
	return
}

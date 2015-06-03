package swfchef

type rectangle struct {
	nbits ubn
	xmin  sbn
	xmax  sbn
	ymin  sbn
	ymax  sbn
}

func (rect *rectangle) Length() ubn {
	return 5 + rect.nbits*4/8 + 1
}

func raw2rectangle(raw []byte) (r rectangle) {
	r = rectangle{}
	r.nbits = raw2ubn(raw, 0, 5)

	n := r.nbits
	r.xmin = raw2sbn(raw, 5+0*n, n)
	r.xmax = raw2sbn(raw, 5+1*n, n)
	r.ymin = raw2sbn(raw, 5+2*n, n)
	r.ymax = raw2sbn(raw, 5+3*n, n)

	return
}

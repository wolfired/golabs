package swfchef

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io/ioutil"
)

const (
	SignatureF byte = 'F'
	SignatureC byte = 'C'
	SignatureZ byte = 'Z'
	SignatureW byte = 'W'
	SignatureS byte = 'S'
)

type Swf struct {
	bs []byte
}

func ReadSwf(path_to_swf string) (r *Swf) {
	bs, err := ioutil.ReadFile(path_to_swf)
	if nil != err {
		return
	}

	r = new(Swf)
	r.bs = bs

	b := bytes.NewBuffer(bs[])
	zr := zlib.NewReader(b)
	zr.Read()

	return
}

func (s *Swf) Size() uint32 {
	return binary.LittleEndian.Uint32(s.bs[4:8])
}

func (s *Swf) Version() uint8 {
	return s.bs[3]
}

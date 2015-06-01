package swfchef

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
	"log"
)

const (
	SignatureF byte = 'F'
	SignatureC byte = 'C'
	SignatureZ byte = 'Z'
	SignatureW byte = 'W'
	SignatureS byte = 'S'

	AttributesUseDirectBlit uint32 = 0x40000000
	AttributesUseGPU        uint32 = 0x20000000
	AttributesHasMetadata   uint32 = 0x10000000
	AttributesActionScript3 uint32 = 0x8000000
	AttributesUseNetwork    uint32 = 0x1000000
)

var (
	swfHandler map[byte]xSwf = map[byte]xSwf{SignatureF: fSwf, SignatureC: cSwf, SignatureZ: zSwf}
)

type xSwf func([]byte) *Swf

func fSwf(bs []byte) *Swf {
	swf := new(Swf)
	swf.bs = bs
	return swf
}

func cSwf(bs []byte) *Swf {
	zr, err := zlib.NewReader(bytes.NewReader(bs[8:]))
	defer zr.Close()
	if nil != err {
		log.Fatal(err)
	}

	size := b2ui32(bs[4:8])

	swf := new(Swf)
	swf.bs = make([]byte, size)

	for i, v := range bs[:8] {
		swf.bs[i] = v
	}

	_, err = zr.Read(swf.bs[8:])
	if nil != err {
		log.Fatal(err)
	}

	return swf
}

func zSwf(bs []byte) *Swf {
	return nil
}

type Swf struct {
	bs []byte
}

func ReadSwf(swffile string) (swf *Swf) {
	bs, err := ioutil.ReadFile(swffile)
	if nil != err {
		log.Fatal(err)
	}
	return swfHandler[bs[0]](bs)
}

func (swf *Swf) Signature() string {
	return string(swf.bs[:3])
}

func (swf *Swf) Version() ui8 {
	return b2ui8(swf.bs[3:4])
}

func (swf *Swf) Length() ui32 {
	return b2ui32(swf.bs[4:8])
}

func (swf *Swf) FrameSize() rectangle {
	return rectangle{}
}

func (swf *Swf) Tag() {

}

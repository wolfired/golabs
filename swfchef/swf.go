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

func fSwf(raw []byte) *Swf {
	swf := new(Swf)
	swf.raw = raw
	return swf
}

func cSwf(raw []byte) *Swf {
	zr, err := zlib.NewReader(bytes.NewReader(raw[8:]))
	defer zr.Close()
	if nil != err {
		log.Fatal(err)
	}

	size := raw2ui32(raw[4:8])

	swf := new(Swf)
	swf.raw = make([]byte, size)

	for i, v := range raw[:8] {
		swf.raw[i] = v
	}

	_, err = zr.Read(swf.raw[8:])
	if nil != err {
		log.Fatal(err)
	}

	return swf
}

func zSwf(raw []byte) *Swf {
	return nil
}

type Swf struct {
	raw       []byte
	frameSize rectangle
}

func ReadSwf(swffile string) (swf *Swf) {
	raw, err := ioutil.ReadFile(swffile)
	if nil != err {
		log.Fatal(err)
	}

	swf = swfHandler[raw[0]](raw)
	swf.frameSize = raw2rectangle(swf.raw[8:])

	return
}

func (swf *Swf) Signature() string {
	return string(swf.raw[:3])
}

func (swf *Swf) Version() ui8 {
	return raw2ui8(swf.raw[3:4])
}

func (swf *Swf) Length() ui32 {
	return raw2ui32(swf.raw[4:8])
}

func (swf *Swf) FrameSize() rectangle {
	return swf.frameSize
}

func (swf *Swf) FrameRate() fixed16 {
	offset := 8 + swf.frameSize.Length()
	return raw2fixed16(swf.raw[offset : offset+2])
}

func (swf *Swf) FrameCount() ui16 {
	offset := 8 + swf.frameSize.Length() + 2
	return raw2ui16(swf.raw[offset : offset+2])
}

func (swf *Swf) Tags() {

}

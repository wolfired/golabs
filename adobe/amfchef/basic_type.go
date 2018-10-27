package amfchef

import (
	"bytes"
	"encoding/binary"
	"math"
)

type u8amf byte

func (u8 *u8amf) encode(buf *bytes.Buffer) {
	buf.WriteByte(byte(*u8))
}

func (u8 *u8amf) decode(buf *bytes.Buffer) {
	b, _ := buf.ReadByte()
	*u8 = u8amf(b)
}

type u16amf uint16

func (u16 *u16amf) encode(buf *bytes.Buffer) {
	p := make([]byte, 2, 2)
	binary.BigEndian.PutUint16(p, uint16(*u16))
	buf.Write(p)
}

func (u16 *u16amf) decode(buf *bytes.Buffer) {
	*u16 = u16amf(binary.BigEndian.Uint16(buf.Next(2)))
}

type u32amf uint32

func (u32 *u32amf) encode(buf *bytes.Buffer) {
	p := make([]byte, 4, 4)
	binary.BigEndian.PutUint32(p, uint32(*u32))
	buf.Write(p)
}

func (u32 *u32amf) decode(buf *bytes.Buffer) {
	*u32 = u32amf(binary.BigEndian.Uint32(buf.Next(4)))
}

type d64amf float64

func (d64 *d64amf) encode(buf *bytes.Buffer) {
	p := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(p, math.Float64bits(float64(*d64)))
	buf.Write(p)
}

func (d64 *d64amf) decode(buf *bytes.Buffer) {
	*d64 = d64amf(math.Float64frombits(binary.BigEndian.Uint64(buf.Next(8))))
}

type u29amf uint32

func (u29 *u29amf) encode(buf *bytes.Buffer) {
	xxxx := uint32(*u29)

	if 0x7F >= xxxx {
		buf.WriteByte(byte(xxxx))
	} else if 0x3FFF >= xxxx {
		buf.WriteByte(byte(0x80 | xxxx>>7))
		buf.WriteByte(byte(0x7F & xxxx))
	} else if 0x1FFFFF >= xxxx {
		buf.WriteByte(byte(0x80 | xxxx>>14))
		buf.WriteByte(byte(0x80 | xxxx>>7))
		buf.WriteByte(byte(0x7F & xxxx))
	} else if 0x1FFFFFFF >= xxxx {
		buf.WriteByte(byte(0x80 | xxxx>>22))
		buf.WriteByte(byte(0x80 | xxxx>>15))
		buf.WriteByte(byte(0x80 | xxxx>>8))
		buf.WriteByte(byte(xxxx))
	} else {
		panic("range exception")
	}
}

func (u29 *u29amf) decode(buf *bytes.Buffer) {
	x3, x2, x1, x0 := byte(0), byte(0), byte(0), byte(0)

	x0, _ = buf.ReadByte()
	if 0 != x0&0x80 {
		x1 = x0 & 0x7F
		x0, _ = buf.ReadByte()
		if 0 != x0&0x80 {
			x2, x1 = x1, x0&0x7F
			x0, _ = buf.ReadByte()
			if 0 != x0&0x80 {
				x3, x2, x1 = x2, x1, x0&0x7F
				x0, _ = buf.ReadByte()

				*u29 = u29amf(uint32(x3)<<22 + uint32(x2)<<15 + uint32(x1)<<8 + uint32(x0))

				return
			}
		}
	}

	*u29 = u29amf(uint32(x3)<<21 + uint32(x2)<<14 + uint32(x1)<<7 + uint32(x0))
}

type sRef = u29amf
type sValue = u29amf
type utf8empty = bool

type utf8vr struct {
	isRef bool
	ref   sRef
	str   string
}

func (utf8 *utf8vr) isEmpty() bool {
	if utf8.isRef {
		s, _ := stringReferenceTable[utf8.ref]
		return "" == s
	}

	return "" == utf8.str
}

func (utf8 *utf8vr) empty() {
	utf8.isRef = false
	utf8.ref = 0
	utf8.str = ""
}

func (utf8 *utf8vr) setRef(ref sRef) {
	if !utf8.isRef {
		utf8.isRef = true
		utf8.str = ""
	}

	utf8.ref = ref
}

func (utf8 *utf8vr) setStr(str string) {
	if utf8.isRef {
		utf8.isRef = false
		utf8.ref = 0
	}

	utf8.str = str
}

func (utf8 *utf8vr) value() string {
	if utf8.isRef {
		s, _ := stringReferenceTable[utf8.ref]
		return s
	}

	return utf8.str
}

func (utf8 *utf8vr) encode(buf *bytes.Buffer) {
	if utf8.isEmpty() {
		buf.WriteByte(0x01)
		return
	}

	var u29 u29amf

	if utf8.isRef {
		u29 = utf8.ref << 1
		u29.encode(buf)
		return
	}

	u29 = u29amf(len(utf8.str)<<1 + 1)
	u29.encode(buf)

	buf.WriteString(utf8.str)
}

func (utf8 *utf8vr) decode(buf *bytes.Buffer) {
	var u29 u29amf
	u29.decode(buf)

	u32 := uint32(u29)

	if 0x1 == u32 {
		return
	}

	if 0 == 0x01&u32 {
		utf8.isRef = true
		utf8.ref = u29amf(u32 >> 1)
		return
	}

	utf8.isRef = false
	utf8.str = string(buf.Next(int(u32 >> 1)))
}

type xRef = u29amf
type xValue = u29amf

type xmlvr struct {
	isRef bool
	ref   sRef
	str   string
}

func (xml *xmlvr) encode(buf *bytes.Buffer) {
}

func (xml *xmlvr) decode(buf *bytes.Buffer) {
}

type dRef = u29amf
type dValue = u29amf

type datevr struct {
	isRef bool
	ref   sRef
	str   string
}

func (date *datevr) encode(buf *bytes.Buffer) {
}

func (date *datevr) decode(buf *bytes.Buffer) {
}

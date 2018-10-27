package amfchef

import (
	"bytes"
)

type undefinedType struct{}

func (ut *undefinedType) encode(buf bytes.Buffer) {
	buf.WriteByte(byte(undefinedMarker))
}

func (ut *undefinedType) decode(buf bytes.Buffer) {
	buf.ReadByte()
}

type nullType struct{}

func (nt *nullType) encode(buf bytes.Buffer) {
	buf.WriteByte(byte(nullMarker))
}

func (nt *nullType) decode(buf bytes.Buffer) {
	buf.ReadByte()
}

type falseType bool

func (ft *falseType) encode(buf bytes.Buffer) {
	buf.WriteByte(byte(falseMarker))
}

func (ft *falseType) decode(buf bytes.Buffer) {
	buf.ReadByte()
	*ft = false
}

type trueType bool

func (tt *trueType) encode(buf bytes.Buffer) {
	buf.WriteByte(byte(trueMarker))
}

func (tt *trueType) decode(buf bytes.Buffer) {
	buf.ReadByte()
	*tt = true
}

type integerType u29amf

func (it *integerType) encode(buf *bytes.Buffer) {
	buf.WriteByte(byte(integerMarker))
	(*u29amf)(it).encode(buf)
}

func (it *integerType) decode(buf *bytes.Buffer) {
	buf.ReadByte()
	(*u29amf)(it).decode(buf)
}

type doubleType d64amf

func (dt *doubleType) encode(buf *bytes.Buffer) {
	buf.WriteByte(byte(doubleMarker))
	(*d64amf)(dt).encode(buf)
}

func (dt *doubleType) decode(buf *bytes.Buffer) {
	buf.ReadByte()
	(*d64amf)(dt).decode(buf)
}

type stringType struct {
	utf8vr
}

func (st *stringType) encode(buf *bytes.Buffer) {
	buf.WriteByte(byte(stringMarker))
	st.utf8vr.encode(buf)
}

func (st *stringType) decode(buf *bytes.Buffer) {
	buf.ReadByte()
	st.utf8vr.decode(buf)
}

type xmlDocType struct {
	xmlvr
}

func (xdt *xmlDocType) encode(buf *bytes.Buffer) {
	buf.WriteByte(byte(xmlDocMarker))
	xdt.xmlvr.encode(buf)
}

func (xdt *xmlDocType) decode(buf *bytes.Buffer) {
	buf.ReadByte()
	xdt.xmlvr.decode(buf)
}

package wabf

import (
	"github.com/wolfired/golabs/leb128"
)

type varuint1 uint8

func (v *varuint1) decode(raw []byte) uint {
	r, c := leb128.DecodeU(raw)
	*v = varuint1(r)
	return c
}

type varuint7 uint8

func (v *varuint7) decode(raw []byte) uint {
	r, c := leb128.DecodeU(raw)
	*v = varuint7(r)
	return c
}

type varuint32 uint32

func (v *varuint32) decode(raw []byte) uint {
	r, c := leb128.DecodeU(raw)
	*v = varuint32(r)
	return c
}

type varint7 int8

func (v *varint7) decode(raw []byte) uint {
	r, c := leb128.DecodeS(raw)
	*v = varint7(r)
	return c
}

type varint32 int32

func (v *varint32) decode(raw []byte) uint {
	r, c := leb128.DecodeS(raw)
	*v = varint32(r)
	return c
}

type varint64 int64

func (v *varint64) decode(raw []byte) uint {
	r, c := leb128.DecodeS(raw)
	*v = varint64(r)
	return c
}

type valueType varint7

func (v *valueType) decode(raw []byte) uint {
	r, c := leb128.DecodeS(raw)
	*v = valueType(r)
	return c
}

type blockType varint7

func (b *blockType) decode(raw []byte) uint {
	r, c := leb128.DecodeS(raw)
	*b = blockType(r)
	return c
}

type elemType varint7

func (e *elemType) decode(raw []byte) uint {
	r, c := leb128.DecodeS(raw)
	*e = elemType(r)
	return c
}

type funcType struct {
	form        varint7
	paramCount  varuint32
	paramTypes  []valueType
	returnCount varuint1
	returnType  valueType
}

func (f *funcType) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = f.form.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = f.paramCount.decode(raw)
	raw = raw[used:]
	totalUsed += used

	f.paramTypes = make([]valueType, f.paramCount)
	for i := uint(0); i < uint(f.paramCount); i++ {
		used = f.paramTypes[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	used = f.returnCount.decode(raw)
	raw = raw[used:]
	totalUsed += used

	if 0 < f.returnCount {
		used = f.returnType.decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type globalType struct {
	contentType valueType
	mutability  varuint1
}

func (g *globalType) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = g.contentType.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = g.mutability.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type tableType struct {
	elementType elemType
	limits      resizableLimits
}

func (t *tableType) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = t.elementType.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = t.limits.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type memoryType struct {
	limits resizableLimits
}

func (m *memoryType) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = m.limits.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type externalKind uint8

func (e *externalKind) decode(raw []byte) uint {
	*e = externalKind(raw[0])
	raw = raw[1:]
	return 1
}

const (
	externalKindFunction externalKind = 0
	externalKindTable    externalKind = 1
	externalKindMemory   externalKind = 2
	externalKindGlobal   externalKind = 3
)

type resizableLimits struct {
	flags   varuint1
	initial varuint32
	maximum varuint32
}

func (r *resizableLimits) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = r.flags.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = r.initial.decode(raw)
	raw = raw[used:]
	totalUsed += used

	if 1 == r.flags {
		used = r.maximum.decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

const (
	expressionBodyEnd = byte(opcode_end)
)

type initExpr struct {
	expr []byte
}

func (ie *initExpr) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	for used = 0; true; used++ {
		if expressionBodyEnd == raw[used] {
			used++
			ie.expr = raw[:used]
			raw = raw[used:]
			totalUsed += used
			break
		}
	}

	return totalUsed
}

type importEntry struct {
	moduleLen  varuint32
	moduleStr  string
	fieldLen   varuint32
	fieldStr   string
	kind       externalKind
	importType decoder
}

func (ie *importEntry) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = ie.moduleLen.decode(raw)
	raw = raw[used:]
	totalUsed += used

	ie.moduleStr = string(raw[:ie.moduleLen])
	raw = raw[ie.moduleLen:]
	totalUsed += uint(ie.moduleLen)

	used = ie.fieldLen.decode(raw)
	raw = raw[used:]
	totalUsed += used

	ie.fieldStr = string(raw[:ie.fieldLen])
	raw = raw[ie.fieldLen:]
	totalUsed += uint(ie.fieldLen)

	used = ie.kind.decode(raw)
	raw = raw[used:]
	totalUsed += used

	switch ie.kind {
	case externalKindFunction:
		ie.importType = new(varuint32)
	case externalKindTable:
		ie.importType = &tableType{}
	case externalKindMemory:
		ie.importType = &memoryType{}
	case externalKindGlobal:
		ie.importType = &globalType{}
	}

	used = ie.importType.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type globalVariable struct {
	gtype globalType
	init  initExpr
}

func (gv *globalVariable) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = gv.gtype.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = gv.init.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type exportEntry struct {
	fieldLen varuint32
	fieldStr string
	kind     externalKind
	index    varuint32
}

func (ee *exportEntry) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = ee.fieldLen.decode(raw)
	raw = raw[used:]
	totalUsed += used

	ee.fieldStr = string(raw[:ee.fieldLen])
	raw = raw[ee.fieldLen:]
	totalUsed += uint(ee.fieldLen)

	used = ee.kind.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = ee.index.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type elemSegment struct {
	index   varuint32
	offset  initExpr
	numElem varuint32
	elems   []varuint32
}

func (es *elemSegment) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = es.index.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = es.offset.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = es.numElem.decode(raw)
	raw = raw[used:]
	totalUsed += used

	es.elems = make([]varuint32, es.numElem)

	for i := uint(0); i < uint(es.numElem); i++ {
		used = es.elems[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type functionBody struct {
	bodySize   varuint32
	localCount varuint32
	locals     []localEntry
	code       []byte
}

const (
	functionBodyEnd = byte(opcode_end)
)

func (fb *functionBody) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = fb.bodySize.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = fb.localCount.decode(raw)
	raw = raw[used:]
	totalUsed += used

	fb.locals = make([]localEntry, fb.localCount)

	for i := uint(0); i < uint(fb.localCount); i++ {
		used = fb.locals[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	for used = 0; true; used++ {
		if functionBodyEnd == raw[used] {
			used++
			fb.code = raw[:used]
			raw = raw[used:]
			totalUsed += used
			break
		}
	}

	return totalUsed
}

type localEntry struct {
	count varuint32
	ltype valueType
}

func (le *localEntry) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = le.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = le.ltype.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type dataSegment struct {
	index  varuint32
	offset initExpr
	size   varuint32
	data   []byte
}

func (ds *dataSegment) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = ds.index.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = ds.offset.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = ds.size.decode(raw)
	raw = raw[used:]
	totalUsed += used

	ds.data = raw[:ds.size]
	raw = raw[ds.size:]
	totalUsed += uint(ds.size)

	return totalUsed
}

type nameTypeModule struct {
	nameLen varuint32
	nameStr string
}

func (ntm *nameTypeModule) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = ntm.nameLen.decode(raw)
	raw = raw[used:]
	totalUsed += used

	ntm.nameStr = string(raw[:ntm.nameLen])
	raw = raw[ntm.nameLen:]
	totalUsed += uint(ntm.nameLen)

	return totalUsed
}

type nameMap struct {
	count varuint32
	names []naming
}

func (nm *nameMap) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = nm.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	nm.names = make([]naming, nm.count)

	for i := uint(0); i < uint(nm.count); i++ {
		used = nm.names[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type naming struct {
	index   varuint32
	nameLen varuint32
	nameStr string
}

func (n *naming) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = n.index.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = n.nameLen.decode(raw)
	raw = raw[used:]
	totalUsed += used

	n.nameStr = string(raw[:n.nameLen])
	raw = raw[n.nameLen:]
	totalUsed += uint(n.nameLen)

	return totalUsed
}

type nameTypeFunction = nameMap

type nameTypeLocal struct {
	count varuint32
	funcs []localName
}

func (ntl *nameTypeLocal) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = ntl.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	ntl.funcs = make([]localName, ntl.count)

	for i := uint(0); i < uint(ntl.count); i++ {
		used = ntl.funcs[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type localName struct {
	index    varuint32
	localMap nameMap
}

func (ln *localName) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = ln.index.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = ln.localMap.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

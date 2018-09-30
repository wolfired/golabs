package wabf

const (
	sectionCodeCustom   = 0
	sectionCodeType     = 1
	sectionCodeImport   = 2
	sectionCodeFunction = 3
	sectionCodeTable    = 4
	sectionCodeMemory   = 5
	sectionCodeGlobal   = 6
	sectionCodeExport   = 7
	sectionCodeStart    = 8
	sectionCodeElement  = 9
	sectionCodeCode     = 10
	sectionCodeData     = 11
)

type sectionType struct {
	count   varuint32
	entries []funcType
}

func (st *sectionType) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = st.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	st.entries = make([]funcType, st.count)

	for i := uint(0); i < uint(st.count); i++ {
		used := st.entries[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionImport struct {
	count   varuint32
	entries []importEntry
}

func (si *sectionImport) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = si.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	si.entries = make([]importEntry, si.count)

	for i := uint(0); i < uint(si.count); i++ {
		used = si.entries[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionFunction struct {
	count varuint32
	types []varuint32
}

func (sf *sectionFunction) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = sf.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	sf.types = make([]varuint32, sf.count)

	for i := uint(0); i < uint(sf.count); i++ {
		used = sf.types[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionTable struct {
	count   varuint32
	entries []tableType
}

func (st *sectionTable) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = st.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	st.entries = make([]tableType, st.count)

	for i := uint(0); i < uint(st.count); i++ {
		used = st.entries[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionMemory struct {
	count   varuint32
	entries []memoryType
}

func (sm *sectionMemory) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = sm.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	sm.entries = make([]memoryType, sm.count)

	for i := uint(0); i < uint(sm.count); i++ {
		used = sm.entries[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionGlobal struct {
	count   varuint32
	globals []globalVariable
}

func (sg *sectionGlobal) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = sg.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	sg.globals = make([]globalVariable, sg.count)

	for i := uint(0); i < uint(sg.count); i++ {
		used = sg.globals[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionExport struct {
	count   varuint32
	entries []exportEntry
}

func (se *sectionExport) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = se.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	se.entries = make([]exportEntry, se.count)

	for i := uint(0); i < uint(se.count); i++ {
		used = se.entries[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionStart struct {
	index varuint32
}

func (ss *sectionStart) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = ss.index.decode(raw)
	raw = raw[used:]
	totalUsed += used

	return totalUsed
}

type sectionElement struct {
	count   varuint32
	entries []elemSegment
}

func (se *sectionElement) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = se.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	se.entries = make([]elemSegment, se.count)

	for i := uint(0); i < uint(se.count); i++ {
		used = se.entries[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionCode struct {
	count  varuint32
	bodies []functionBody
}

func (sc *sectionCode) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = sc.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	sc.bodies = make([]functionBody, sc.count)

	for i := uint(0); i < uint(sc.count); i++ {
		used = sc.bodies[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionData struct {
	count   varuint32
	entries []dataSegment
}

func (sd *sectionData) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = sd.count.decode(raw)
	raw = raw[used:]
	totalUsed += used

	sd.entries = make([]dataSegment, sd.count)

	for i := 0; i < len(sd.entries); i++ {
		used = sd.entries[i].decode(raw)
		raw = raw[used:]
		totalUsed += used
	}

	return totalUsed
}

type sectionName struct {
	nameType        varuint7
	namePayloadLen  varuint32
	namepayloadData []byte

	namepayloadType decoder
}

func (sn *sectionName) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = sn.nameType.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = sn.namePayloadLen.decode(raw)
	raw = raw[used:]
	totalUsed += used

	sn.namepayloadData = raw[:sn.namePayloadLen]
	totalUsed += uint(sn.namePayloadLen)

	switch sn.nameType {
	case 0:
		sn.namepayloadType = &nameTypeModule{}
	case 1:
		sn.namepayloadType = &nameTypeFunction{}
	case 2:
		sn.namepayloadType = &nameTypeLocal{}
	}

	if nil != sn.namepayloadType {
		sn.namepayloadType.decode(sn.namepayloadData)
	}

	return totalUsed
}

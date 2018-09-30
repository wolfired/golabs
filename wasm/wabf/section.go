package wabf

type section struct {
	id          varuint7
	payloadLen  varuint32
	nameLen     varuint32
	name        string
	payloadData []byte

	payloadType decoder
}

func (s *section) decode(raw []byte) uint {
	totalUsed, used := uint(0), uint(0)

	used = s.id.decode(raw)
	raw = raw[used:]
	totalUsed += used

	used = s.payloadLen.decode(raw)
	raw = raw[used:]
	totalUsed += used

	if 0 == s.id {
		used = s.nameLen.decode(raw)
		raw = raw[used:]
		totalUsed += used

		s.name = string(raw[:s.nameLen])
		raw = raw[s.nameLen:]
		totalUsed += uint(s.nameLen)

		s.payloadData = raw[:uint(s.payloadLen)-used-uint(s.nameLen)]
	} else {
		s.payloadData = raw[:s.payloadLen]
	}
	raw = raw[len(s.payloadData):]
	totalUsed += uint(len(s.payloadData))

	switch s.id {
	case 0:
		s.payloadType = &sectionName{}
	case 1:
		s.payloadType = &sectionType{}
	case 2:
		s.payloadType = &sectionImport{}
	case 3:
		s.payloadType = &sectionFunction{}
	case 4:
		s.payloadType = &sectionTable{}
	case 5:
		s.payloadType = &sectionMemory{}
	case 6:
		s.payloadType = &sectionGlobal{}
	case 7:
		s.payloadType = &sectionExport{}
	case 8:
		s.payloadType = &sectionStart{}
	case 9:
		s.payloadType = &sectionElement{}
	case 10:
		// s.payloadType = &sectionCode{}
	case 11:
		s.payloadType = &sectionData{}
	}

	if nil != s.payloadType {
		s.payloadType.decode(s.payloadData)
	}

	return totalUsed
}

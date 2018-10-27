package amfchef

import "bytes"

type codec interface {
	encode(buf *bytes.Buffer)
	decode(buf *bytes.Buffer)
}

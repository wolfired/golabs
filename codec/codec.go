package codec

import (
	"bytes"
)

/*
Encoder 编码器
*/
type Encoder interface {
	Encode(buf bytes.Buffer)
}

/*
Decoder 解码器
*/
type Decoder interface {
	Decode(buf bytes.Buffer)
}

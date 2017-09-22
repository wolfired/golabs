package bitorent

import (
	"bytes"
	"strconv"

	"github.com/wolfired/golabs/auto"
)

func encodei(i uint64, buf *bytes.Buffer) {
	buf.WriteString("i" + strconv.FormatUint(i, 10) + "e")
}

func encodes(s string, buf *bytes.Buffer) {
	buf.WriteString(strconv.FormatUint(uint64(len([]rune(s))), 10) + ":" + s)
}

func encodel(l []interface{}, buf *bytes.Buffer) {
	buf.WriteString("l")

	for _, v := range l {
		switch v.(type) {
		case int:
			encodei(v.(uint64), buf)
		case string:
			encodes(v.(string), buf)
		case []interface{}:
			encodel(v.([]interface{}), buf)
		case map[string]interface{}:
			encoded(v.(map[string]interface{}), buf)
		}
	}

	buf.WriteString("e")
}

func encoded(d map[string]interface{}, buf *bytes.Buffer) {
	buf.WriteString("d")

	for k, v := range d {
		encodes(k, buf)

		switch v.(type) {
		case int:
			encodei(v.(uint64), buf)
		case string:
			encodes(v.(string), buf)
		case []interface{}:
			encodel(v.([]interface{}), buf)
		case map[interface{}]interface{}:
			encoded(v.(map[string]interface{}), buf)
		}
	}

	buf.WriteString("e")
}

func decodei(buf *bytes.Buffer) uint64 {
	buf.ReadRune()

	bs, _ := buf.ReadBytes('e')
	i, _ := strconv.ParseUint(string(bs[0:len(bs)-1]), 10, 0)

	return i
}

func decodes(buf *bytes.Buffer) string {
	bs, _ := buf.ReadBytes(':')
	c, _ := strconv.Atoi(string(bs[0 : len(bs)-1]))
	s := string(buf.Next(c))

	return s
}

func decodel(buf *bytes.Buffer) []interface{} {
	buf.ReadRune()

	l := make([]interface{}, 0, 8)

	for k, _, _ := buf.ReadRune(); 'e' != k; k, _, _ = buf.ReadRune() {
		buf.UnreadRune()

		switch {
		case 'i' == k:
			l = append(l, decodei(buf))
		case '0' <= k && k <= '9':
			l = append(l, decodes(buf))
		case 'l' == k:
			l = append(l, decodel(buf))
		case 'd' == k:
			l = append(l, decoded(buf))
		}
	}

	return l
}

func decoded(buf *bytes.Buffer) map[string]interface{} {
	buf.ReadRune()

	d := make(map[string]interface{})

	for m, _, _ := buf.ReadRune(); 'e' != m; m, _, _ = buf.ReadRune() {
		buf.UnreadRune()

		k := decodes(buf)

		m, _, _ = buf.ReadRune()
		buf.UnreadRune()

		switch {
		case 'i' == m:
			d[k] = decodei(buf)
		case '0' <= m && m <= '9':
			d[k] = decodes(buf)
		case 'l' == m:
			d[k] = decodel(buf)
		case 'd' == m:
			d[k] = decoded(buf)
		}
	}

	return d
}

/*Decode 解码*/
func Decode(bs []byte, mi interface{}) {
	buf := bytes.NewBuffer(bs)

	d := decoded(buf)

	auto.FillStruct(mi, d, tagName)
}

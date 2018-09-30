package wabf

import (
	"io/ioutil"
	"testing"
)

func Test_decode(t *testing.T) {
	bs, _ := ioutil.ReadFile("./test.wasm")
	m := module{}
	m.decode(bs)
	t.Fatal(m.sections[3].payloadType)
	// t.Fatal(m.sections[0].payloadType)
	// t.Fatal(m.sections[1])
}

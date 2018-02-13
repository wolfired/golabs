package bitorent

import "testing"
import "bytes"
import "fmt"

func Test_encodei(t *testing.T) {
	bs := make([]byte, 1024)
	buf := bytes.NewBuffer(bs)

	buf.Reset()
	encodei(23, buf)
	encodei(-23, buf)
	encodei(0, buf)
	but := string(buf.Bytes())
	expect := "i23ei-23ei0e"
	if expect != but {
		t.Error("encodei error") 
	}
}

func Test_encodes(t *testing.T) {
	bs := make([]byte, 1024)
	buf := bytes.NewBuffer(bs)

	buf.Reset()
	encodes("23", buf)
	encodes("abc", buf)
	encodes("", buf)
	but := string(buf.Bytes())
	expect := "2:233:abc0:"
	if expect != but {
		t.Error("encodes error")
	}
}

func Test_encodel(t *testing.T) {
	bs := make([]byte, 1024)
	buf := bytes.NewBuffer(bs)

	buf.Reset()
	is := []interface{}{1, 3, 4, 6}
	encodel(is, buf)
	but := string(buf.Bytes())
	expect := "li1ei3ei4ei6ee"
	if expect != but {
		t.Error("encodes error")
	}

	buf.Reset()
	is = []interface{}{"aa", "bbc", "ddd", "aas3"}
	encodel(is, buf)
	but = string(buf.Bytes())
	expect = "l2:aa3:bbc3:ddd4:aas3e"
	if expect != but {
		t.Error("encodes error")
	}

	buf.Reset()
	is = []interface{}{[]interface{}{1}}
	encodel(is, buf)
	but = string(buf.Bytes())
	expect = "lli1eee"
	if expect != but {
		t.Error("encodes error")
	}

	buf.Reset()
	is = []interface{}{map[string]interface{}{"aa": "bb"}}
	encodel(is, buf)
	but = string(buf.Bytes())
	expect = "ld2:aa2:bbee"
	if expect != but {
		t.Error("encodes error")
	}

	fmt.Println(buf)
}

func Test_encoded(t *testing.T) {
	bs := make([]byte, 1024)
	buf := bytes.NewBuffer(bs)

	buf.Reset()
	is := map[string]interface{}{"a": 1}
	encoded(is, buf)
	but := string(buf.Bytes())
	expect := "d1:ai1ee"
	if expect != but {
		t.Error("encodes error")
	}

	buf.Reset()
	is = map[string]interface{}{"a": "b"}
	encoded(is, buf)
	but = string(buf.Bytes())
	expect = "d1:a1:be"
	if expect != but {
		t.Error("encodes error")
	}

	buf.Reset()
	is = map[string]interface{}{"a": []interface{}{1,2}}
	encoded(is, buf)
	but = string(buf.Bytes())
	expect = "d1:ali1ei2eee"
	if expect != but {
		t.Error("encodes error")
	}

	buf.Reset()
	is = map[string]interface{}{"a": map[string]interface{}{"b": "c"}}
	encoded(is, buf)
	but = string(buf.Bytes())
	expect = "d1:ad1:b1:cee"
	if expect != but {
		t.Error("encodes error")
	}

	fmt.Println(buf)
}

func Test_decodei(t *testing.T) {
	buf := bytes.NewBufferString("i3ei88e")

	but := decodei(buf)
	expect := 3
	if expect != but {
		t.Error("encodes error")
	}

	but = decodei(buf)
	expect = 88
	if expect != but {
		t.Error("encodes error")
	}
}

func Test_decodes(t *testing.T) {
	buf := bytes.NewBufferString("0:1:a2:bb")

	but := decodes(buf)
	expect := ""
	if expect != but {
		t.Error("encodes error")
	}

	but = decodes(buf)
	expect = "a"
	if expect != but {
		t.Error("encodes error")
	}

	but = decodes(buf)
	expect = "bb"
	if expect != but {
		t.Error("encodes error")
	}
}

func Test_decodel(t *testing.T) {
	buf := bytes.NewBufferString("ld1:al1:aeee")
	l := decodel(buf)
	fmt.Println(l)
}

func Test_decoded(t *testing.T) {
	buf := bytes.NewBufferString("d1:ald1:al1:aeeee")
	d := decoded(buf)
	fmt.Println(d)
}

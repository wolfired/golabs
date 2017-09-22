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

	expect := "i23ei-23ei0e"
	but := string(buf.Bytes())

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

	expect := "2:233:abc0:"
	but := string(buf.Bytes())

	if expect != but {
		t.Error("encodes error")
	}
}

func Test_encodel(t *testing.T) {
	bs := make([]byte, 1024)

	buf := bytes.NewBuffer(bs)
	buf.Reset()

	// is := []interface{}{1, 3, 4, 6}
	// encodel(is, buf)

	// is = []interface{}{"aa", "bbc", "ddd", "aas3"}
	// encodel(is, buf)

	// is := []interface{}{[]interface{}{1}}
	// encodel(is, buf)

	is := []interface{}{map[string]interface{}{"aa": "bb"}}
	encodel(is, buf)

	fmt.Println(buf)
}

func Test_decodei(t *testing.T) {
	buf := bytes.NewBufferString("i3ei88e")
	i := decodei(buf)
	fmt.Println(i)
}

func Test_decodes(t *testing.T) {
	buf := bytes.NewBufferString("0:")
	s := decodes(buf)
	fmt.Println(s)
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

package amfchef

import (
	"bytes"
	"testing"
)

func Test_integerType(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	var it integerType

	it = 321568    //2**29 - 1
	it.encode(buf) //[]byte{4, 0xFF, 0xFF, 0xFF, 0xFF}

	it = 6666665
	it.encode(buf) //[]byte{4, 0}

	it.decode(buf)
	if 321568 != it {
		t.Error("expect: ", 321568, ", but: ", it)
	}

	it.decode(buf)
	if 6666665 != it {
		t.Error("expect: ", 6666665, ", but: ", it)
	}
}

func Test_doubleType(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	var dt doubleType

	dt = 1.236
	dt.encode(buf) //[]byte{5, 63, 243, 198, 167, 239, 157, 178, 45}

	dt = -1.359
	dt.encode(buf) //[]byte{5, 191, 245, 190, 118, 200, 180, 57, 88}

	dt.decode(buf)
	if 1.236 != dt {
		t.Error("expect: ", 1.236, ", but: ", dt)
	}

	dt.decode(buf)
	if -1.359 != dt {
		t.Error("expect: ", -1.359, ", but: ", dt)
	}
}

func Test_stringType(t *testing.T) {
	stringReferenceTable[0] = "China"

	buf := bytes.NewBuffer([]byte{})

	var st stringType

	st.setStr("how")
	st.encode(buf)

	st.decode(buf)
	if "how" != st.value() {
		t.Error("expect: ", "how", ", but: ", st)
	}

	st.setRef(0)
	st.encode(buf)

	st.decode(buf)
	if "China" != st.value() {
		t.Error("expect: ", "China", ", but: ", st)
	}

	st.setStr("UTF-8")
	st.encode(buf)

	st.decode(buf)
	if "UTF-8" != st.value() {
		t.Error("expect: ", "UTF-8", ", but: ", st)
	}

	st.empty()
	st.encode(buf)

	st.decode(buf)
	if "" != st.value() {
		t.Error("expect: ", "", ", but: ", st)
	}
}

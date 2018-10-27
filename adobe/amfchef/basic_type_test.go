package amfchef

import (
	"bytes"
	"math"
	"testing"
)

func Test_d64amf(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	var d64 d64amf

	d64 = math.MaxFloat64
	d64.encode(buf) //[]byte{127, 239, 255, 255, 255, 255, 255, 255}

	d64 = math.SmallestNonzeroFloat64
	d64.encode(buf) //[]byte{0, 0, 0, 0, 0, 0, 0, 1}

	d64 = d64amf(math.Inf(1))
	d64.encode(buf) //[]byte{127, 240, 0, 0, 0, 0, 0, 0}

	d64 = d64amf(math.Inf(-1))
	d64.encode(buf) //[]byte{255, 240, 0, 0, 0, 0, 0, 0}

	d64 = d64amf(math.NaN())
	d64.encode(buf) //[]byte{255, 248, 0, 0, 0, 0, 0, 0}

	d64.decode(buf)
	if math.MaxFloat64 != d64 {
		t.Error("expect: ", math.MaxFloat64, ", but: ", d64)
	}

	d64.decode(buf)
	if math.SmallestNonzeroFloat64 != d64 {
		t.Error("expect: ", math.SmallestNonzeroFloat64, ", but: ", d64)
	}

	d64.decode(buf)
	if !math.IsInf(float64(d64), 1) {
		t.Error("expect: ", math.Inf(1), ", but: ", d64)
	}

	d64.decode(buf)
	if !math.IsInf(float64(d64), -1) {
		t.Error("expect: ", math.Inf(-1), ", but: ", d64)
	}

	d64.decode(buf)
	if !math.IsNaN(float64(d64)) {
		t.Error("expect: ", math.NaN(), ", but: ", d64)
	}
}

func Test_u29amf(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})

	var u29 u29amf

	u29 = 536870911 //2**29 - 1
	u29.encode(buf) //[]byte{0xFF, 0xFF, 0xFF, 0xFF}

	u29 = 0
	u29.encode(buf) //[]byte{0}

	u29.decode(buf)
	if 536870911 != u29 {
		t.Error("expect: ", 536870911, ", but: ", u29)
	}

	u29.decode(buf)
	if 0 != u29 {
		t.Error("expect: ", 0, ", but: ", u29)
	}
}

func Test_utf8vr(t *testing.T) {
	stringReferenceTable[0] = "hello"

	var utf8 *utf8vr

	utf8 = &utf8vr{}

	if !utf8.isEmpty() {
		t.Error("expect: ", true, ", but: ", false)
	}

	utf8.str = "hi"

	if utf8.isEmpty() {
		t.Error("expect: ", true, ", but: ", false)
	}

	if "hi" != utf8.value() {
		t.Error("expect: ", "hi", ", but: ", utf8.value())
	}

	utf8.empty()

	if !utf8.isEmpty() {
		t.Error("expect: ", true, ", but: ", false)
	}

	utf8.setRef(0)

	if utf8.isEmpty() {
		t.Error("expect: ", true, ", but: ", false)
	}

	if "hello" != utf8.value() {
		t.Error("expect: ", "hello", ", but: ", utf8.value())
	}

	utf8.setRef(1)

	if !utf8.isEmpty() {
		t.Error("expect: ", true, ", but: ", false)
	}

	if "" != utf8.value() {
		t.Error("expect: ", "hello", ", but: ", utf8.value())
	}
}

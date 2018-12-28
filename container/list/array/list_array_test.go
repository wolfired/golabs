package array

import (
	"testing"
)

func Test_New(t *testing.T) {
	{
		a := New(0).(*array)
		if nil == a || 0 != len(a.raw) {
			t.Error("But", a)
		}
	}

	{
		a := New(1).(*array)
		if nil == a || 1 != len(a.raw) {
			t.Error("But", a)
		}
	}
}

func Test_Size(t *testing.T) {
	{
		n := 0
		a := New(n)
		if n != a.Size() {
			t.Error("Want", n)
			t.Error("But", a.Size())
		}
	}

	{
		n := 1
		a := New(n)
		if n != a.Size() {
			t.Error("Want", n)
			t.Error("But", a.Size())
		}
	}

	{
		n := 10
		a := New(n)
		if n != a.Size() {
			t.Error("Want", n)
			t.Error("But", a.Size())
		}
	}
}

func Test_Assign(t *testing.T) {
	a := New(0).(*array)
	a.Assign(0, 0)
}

func Test_ValueAt(t *testing.T) {
	a := New(0).(*array)
	a.ValueAt(0)
}

func Test_IndexOf(t *testing.T) {
	a := New(0).(*array)
	i := a.IndexOf(0)
	if -1 != i {
		t.Error("Want", -1)
		t.Error("But", i)
	}
}

func Test_Insert(t *testing.T) {
	a := New(0).(*array)
	a.Insert(0, 0)
}

func Test_shiftR(t *testing.T) {
	{
		n := 0
		a := mockNewArray(n)
		a.shiftR(0, 0)
		but := [0]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 0
		a := mockNewArray(n)
		a.shiftR(0, 1)
		but := [0]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 0
		a := mockNewArray(n)
		a.shiftR(0, 2)
		but := [0]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 1
		a := mockNewArray(n)
		a.shiftR(0, 0)
		but := [1]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{0}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 1
		a := mockNewArray(n)
		a.shiftR(0, 1)
		but := [1]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{nil}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 1
		a := mockNewArray(n)
		a.shiftR(0, 2)
		but := [1]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{nil}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 10
		a := mockNewArray(n)
		a.shiftR(0, 1)
		but := [10]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{nil, 0, 1, 2, 3, 4, 5, 6, 7, 8}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 10
		a := mockNewArray(n)
		a.shiftR(1, 9)
		but := [10]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{0, nil, nil, nil, nil, nil, nil, nil, nil, nil}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}
}

func Test_shiftL(t *testing.T) {
	{
		n := 0
		a := mockNewArray(n)
		a.shiftL(0, 0)
		but := [0]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 0
		a := mockNewArray(n)
		a.shiftL(0, 1)
		but := [0]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 0
		a := mockNewArray(n)
		a.shiftL(0, 2)
		but := [0]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 1
		a := mockNewArray(n)
		a.shiftL(0, 0)
		but := [1]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{0}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 1
		a := mockNewArray(n)
		a.shiftL(0, 1)
		but := [1]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{nil}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 1
		a := mockNewArray(n)
		a.shiftL(0, 2)
		but := [1]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{nil}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 10
		a := mockNewArray(n)
		a.shiftL(0, 1)
		but := [10]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, nil}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}

	{
		n := 10
		a := mockNewArray(n)
		a.shiftL(1, 8)
		but := [10]interface{}{}
		copy(but[:], a.raw)
		want := [...]interface{}{0, 9, nil, nil, nil, nil, nil, nil, nil, nil}
		if want != but {
			t.Error("Want", want)
			t.Error("But", a.raw)
		}
	}
}

func mockNewArray(n int) *array {
	a := New(n).(*array)

	for i := 0; i < n; i++ {
		a.Assign(i, i)
	}

	return a
}

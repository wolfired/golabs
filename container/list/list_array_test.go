package list

import (
	"testing"
)

func Test_Size(t *testing.T) {
	s := 10
	a := NewArray(s)
	if s != a.Size() {
		t.Error("Want", s)
		t.Error("But", a.Size())
	}
}

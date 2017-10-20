package container

import (
	"testing"
)

func Test_IndexAt(t *testing.T) {
	a := newArrayList(10)
	if 10 != a.Size() {
		t.Error("Want", 10)
		t.Error("But", a.Size())
	}
}

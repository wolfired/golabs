package display

import (
	"testing"
)

func Test_Object(t *testing.T) {
	o := CreateObject("one")

	if "two" != o.IEntity.Name() {
		t.Error("But", o.IEntity.Name())
	}

	if nil != o.Parent() {
		t.Error("But", o.Parent())
	}
}

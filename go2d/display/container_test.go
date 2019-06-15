package display

import (
	"testing"
)

func Test_Container(t *testing.T) {
	c := CreateContainer("ONE")

	want := "ONE"
	but := c.Name()

	if want != but {
		t.Error("But", but)
	}
}

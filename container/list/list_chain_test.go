package list

import (
	"testing"
)

func chainInsert(c Lister, vs ...any) {
	for i, v := range vs {
		c.Insert(i, v)
	}
}

func chainDelete(c Lister, is ...int) {
	for _, v := range is {
		c.Delete(v)
	}
}

func Test_NewChain(t *testing.T) {
	n := 3
	c := NewChain()
	chainInsert(c, 0, 1, 2, 3)
	chainDelete(c, 0)
	if n != c.Size() {
		t.Error("Want", n)
		t.Error("But", c.Size())
	}
}

func Test_Pos(t *testing.T) {
	c := NewChain()
	c.Insert(0, 0)
	c.Insert(0, 1)
	c.Insert(0, 2)
	c.Insert(1, 4)
	c.Delete(0)
	if 2 != c.Pos(0) {
		t.Error("Want", 2)
		t.Error("But", c.Pos(0))
	}

	if -1 != c.Pos(99) {
		t.Error("Want", -1)
		t.Error("But", c.Pos(0))
	}
}

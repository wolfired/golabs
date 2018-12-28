package chain

import (
	"testing"
)

func Test_New(t *testing.T) {
	c := New().(*chain)
	if nil == c || 0 != c.size || nil == c.head || nil != c.head.next || nil != c.head.value {
		t.Error("But", c)
	}
}

func Test_Assign(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(r.(error).Error())
		}
	}()

	{
		c := New().(*chain)
		c.Assign(0, 0)
		if 0 != c.ValueAt(0) {
			t.Error("Want", 0)
			t.Error("But", c.ValueAt(0))
		}
	}
}

func Test_ValueAt(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error(r.(error).Error())
		}
	}()

	c := New().(*chain)
	if nil != c.ValueAt(0) {
		t.Error("Want", nil)
		t.Error("But", c.ValueAt(0))
	}
}

func Test_IndexOf(t *testing.T) {
	c := New().(*chain)
	i := c.IndexOf(nil)
	if -1 != i {
		t.Error("Want", -1)
		t.Error("But", i)
	}
}

func Test_Insert(t *testing.T) {
	c := New().(*chain)
	c.Insert(0, 2)
	c.Insert(0, 1)
	c.Insert(0, 0)
	c.print()
	t.Error("Want")
	// if 0 != c.ValueAt(0) {
	// 	t.Error("Want", 0)
	// 	t.Error("But", c.ValueAt(0))
	// }
}

func mockNewChain(n int) *chain {
	c := New().(*chain)

	for i := 0; i < n; i++ {
		c.Insert(i, i)
	}

	return c
}

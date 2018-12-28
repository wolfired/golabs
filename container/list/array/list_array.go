package array

import (
	"github.com/wolfired/golabs/container/list"
)

/*
New 新建一个容量为n的数组, 容量不变
*/
func New(n int) list.Lister {
	return &array{raw: make([]interface{}, n)}
}

type array struct {
	raw []interface{}
}

func (a *array) Size() int {
	return len(a.raw)
}

func (a *array) Assign(i int, v interface{}) { //O(1)
	a.raw[i] = v
}

func (a *array) ValueAt(i int) interface{} { //O(1)
	return a.raw[i]
}

func (a *array) IndexOf(v interface{}) int { //O(n)
	for ri, rv := range a.raw {
		if rv == v {
			return ri
		}
	}
	return -1
}

func (a *array) Insert(i int, v interface{}) { //O(n)
	a.shiftR(i, 1)
	a.Assign(i, v)
}

func (a *array) Delete(i int) interface{} { //O(n)
	v := a.ValueAt(i)
	a.shiftL(i, 1)
	return v
}

func (a *array) shiftR(b int, s int) {
	n := a.Size()
	for i := n - 1; i > b-1; i-- {
		if i > b+s-1 {
			a.raw[i] = a.raw[i-s]
		} else {
			a.raw[i] = nil
		}
	}
}

func (a *array) shiftL(b int, s int) {
	n := a.Size()
	for i := b; i < n; i++ {
		if i+s < n {
			a.raw[i] = a.raw[i+s]
		} else {
			a.raw[i] = nil
		}
	}
}

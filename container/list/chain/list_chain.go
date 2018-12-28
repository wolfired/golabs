package chain

import (
	"fmt"

	"github.com/wolfired/golabs/container/list"
)

/*
New 新建一个容量为0的链表, 容量可变
*/
func New() list.Lister {
	return &chain{head: &node{}, size: 0}
}

type node struct {
	next  *node
	value interface{}
}

type chain struct {
	head *node
	size int
}

func (c *chain) Size() int {
	return c.size
}

func (c *chain) Assign(i int, v interface{}) { //O(n)
	c.nodeAt(i).value = v
}

func (c *chain) ValueAt(i int) interface{} { //O(n)
	return c.nodeAt(i).value
}

func (c *chain) IndexOf(v interface{}) int { //O(n)
	r, i := -1, 0
	c.forward(c.head.next, func(n *node) bool {
		if n.value == v {
			r = i
			return true
		}
		i++
		return false
	})
	return r
}

func (c *chain) Insert(i int, v interface{}) { //O(1)
	n := c.nodePrev(i)
	n.next = &node{next: n.next, value: v}
	c.size++
}

func (c *chain) Delete(i int) interface{} { //O(1)
	n := c.nodePrev(i)
	n.next = n.next.next
	c.size--
	return n.value
}

func (c *chain) nodeAt(i int) (r *node) {
	c.forward(c.head.next, func(n *node) bool {
		if 0 == i {
			r = n
			return true
		}
		i--
		return false
	})
	return
}

func (c *chain) nodePrev(i int) (r *node) {
	c.forward(c.head, func(n *node) bool {
		if 0 == i {
			r = n
			return true
		}
		i--
		return false
	})
	return
}

func (c *chain) forward(n *node, fn func(n *node) bool) {
	for ; nil != n; n = n.next {
		if fn(n) {
			break
		}
	}
}

func (c *chain) print() {
	c.forward(c.head.next, func(n *node) bool {
		fmt.Println(n.value)
		return false
	})
}

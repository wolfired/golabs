package list

func NewChain() Lister {
	return &chain{head: &node{}, size: 0}
}

/*
chain 单向链
*/
type chain struct {
	head *node
	size int
}

func (c *chain) Size() int {
	return c.size
}

func (c *chain) Set(i int, v any) { //O(n)
	c.get(i + 1).value = v
}

func (c *chain) Get(i int) any { //O(n)
	return c.get(i + 1).value
}

func (c *chain) Insert(i int, v any) { //O(1)
	n := c.get(i)
	n.nxt = &node{nxt: n.nxt, value: v}
	c.size++
}

func (c *chain) Delete(i int) { //O(1)
	n := c.get(i)
	n.nxt = n.nxt.nxt
	c.size--
}

func (c *chain) Pos(v any) int { //O(n)
	for i, n := 0, c.head.nxt; nil != n; i, n = i+1, n.nxt {
		if n.value == v {
			return i
		}
	}
	return -1
}

func (c *chain) get(i int) *node {
	n := c.head
	for ; 0 < i; i-- {
		n = n.nxt
	}
	return n
}

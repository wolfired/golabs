package container

func NewChain(n int) Lister {
	return &chain{}
}

/*
chain 链
*/
type chain struct {
}

func (c *chain) Size() int {
	return 0
}

func (c *chain) Set(i int, v any) {
}

func (c *chain) Get(i int) any {
	return nil
}

func (c *chain) Pos(e any) int {
	return 0
}

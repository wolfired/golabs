package container

func NewArray(n int) Lister {
	return &array{s: make([]any, n)}
}

/*
array 数组
*/
type array struct {
	s []any
}

func (a *array) Size() int {
	return len(a.s)
}

func (a *array) Set(i int, v any) { //0(1)
	a.s[i] = v
}

func (a *array) Get(i int) any { //0(1)
	return a.s[i]
}

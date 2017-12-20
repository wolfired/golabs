package list

func NewArray(n int) Lister {
	return &array{raw: make([]any, n)}
}

/*
array 数组
*/
type array struct {
	raw []any
}

func (a *array) Size() int {
	return len(a.raw)
}

func (a *array) Set(i int, v any) { //O(1)
	a.raw[i] = v
}

func (a *array) Get(i int) any { //O(1)
	return a.raw[i]
}

func (a *array) Insert(i int, v any) { //O(n)
	a.shiftR(i, 1)
	a.Set(i, v)
}

func (a *array) Delete(i int) { //O(n)
	a.shiftL(i, 1)
}

func (a *array) Pos(v any) int { //O(n)
	for i, v_ := range a.raw {
		if v_ == v {
			return i
		}
	}
	return -1
}

func (a *array) shiftR(b int, s int) {
	len := a.Size()
	for i := len - 1; i > b-1; i-- {
		if i > b+s-1 {
			a.raw[i] = a.raw[i-s]
		} else {
			a.raw[i] = nil
		}
	}
}

func (a *array) shiftL(b int, s int) {
	len := a.Size()
	for i := b; i < len; i++ {
		if i+s < len {
			a.raw[i] = a.raw[i+s]
		} else {
			a.raw[i] = nil
		}
	}
}

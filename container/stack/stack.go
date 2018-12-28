package stack

import (
	"github.com/wolfired/golabs/container"
	"github.com/wolfired/golabs/container/list"
)

/*
Stacker 栈
*/
type Stacker interface {
	container.Sizer
	container.Pusher
	container.Popper
}

/*
New 新建一个栈
*/
func New(l list.Lister) Stacker {
	return &stack{Lister: l}
}

type stack struct {
	list.Lister
}

func (s *stack) Size() int {
	return s.Lister.Size()
}

func (s *stack) Push(v interface{}) {
	s.Insert(s.Size()-1, v)
}

func (s *stack) Pop() interface{} {
	e := s.Size() - 1
	v := s.ValueAt(e)
	s.Delete(e)
	return v
}

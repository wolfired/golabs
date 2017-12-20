package stack

import (
	. "github.com/wolfired/golabs/container"
	. "github.com/wolfired/golabs/container/list"
)

type any = interface{}

/*
Stacker 栈
*/
type Stacker interface {
	Sizer
	Pusher
	Popper
}

func NewStack() Stacker {
	return &stack{Lister: NewChain()}
}

type stack struct {
	Lister
}

func (s *stack) Size() int {
	return s.Lister.Size()
}

func (s *stack) Push(v any) {
	s.Insert(s.Size(), v)
}

func (s *stack) Pop() any {
	e := s.Size() - 1
	v := s.Get(e)
	s.Delete(e)
	return v
}

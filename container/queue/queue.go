package queue

import (
	. "github.com/wolfired/golabs/container"
	. "github.com/wolfired/golabs/container/list"
)

type any = interface{}

/*
Queuer 队列
*/
type Queuer interface {
	Sizer
	Pusher
	Puller
}

func NewQueue() Queuer {
	return &queue{Lister: NewChain()}
}

type queue struct {
	Lister
}

func (q *queue) Size() int {
	return q.Lister.Size()
}

func (q *queue) Push(v any) {
	q.Insert(q.Size(), v)
}

func (q *queue) Pull() any {
	v := q.Get(0)
	q.Delete(0)
	return v
}

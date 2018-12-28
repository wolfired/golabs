package queue

import (
	"github.com/wolfired/golabs/container"
	"github.com/wolfired/golabs/container/list"
)

/*
Queuer 队列
*/
type Queuer interface {
	container.Sizer
	container.Pusher
	container.Puller
}

/*
New 新建一个队列
*/
func New(l list.Lister) Queuer {
	return &queue{Lister: l}
}

type queue struct {
	list.Lister
}

func (q *queue) Size() int {
	return q.Lister.Size()
}

func (q *queue) Push(v interface{}) {
	q.Insert(q.Size(), v)
}

func (q *queue) Pull() interface{} {
	v := q.ValueAt(0)
	q.Delete(0)
	return v
}

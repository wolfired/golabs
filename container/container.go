package container

import "errors"

var (
	/*ErrIndexOutOfRange 索引越界 */
	ErrIndexOutOfRange = errors.New("索引越界")
)

/*
Sizer 长度
*/
type Sizer interface {
	Size() int
}

/*
Inserter 插入

0 1 3 4 5
    ↑2
0 1 2 3 4 5
*/
type Inserter interface {
	Insert(i int, v interface{})
}

/*
Deleter 删除

0 1 2 3 4 5
    x
0 1 3 4 5
*/
type Deleter interface {
	Delete(i int) interface{}
}

/*
Pusher 压入

0 1 2 3 4   5←
0 1 2 3 4 5
*/
type Pusher interface {
	Push(v interface{})
}

/*
Popper 弹出

0 1 2 3 4 5→
0 1 2 3 4
*/
type Popper interface {
	Pop() interface{}
}

/*
Puller 拉出

←0 1 2 3 4 5
   1 2 3 4 5
*/
type Puller interface {
	Pull() interface{}
}

/*
Indexer 索引
*/
type Indexer interface {
	Assign(i int, v interface{})
	ValueAt(i int) interface{}
	IndexOf(v interface{}) int
}

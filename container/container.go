package container

type any = interface{}

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
	Insert(i int, v any)
}

/*
Deleter 删除

0 1 2 3 4 5
    x2
0 1 3 4 5
*/
type Deleter interface {
	Delete(i int)
}

/*
Pusher 压入

0 1 2 3 4   5←
0 1 2 3 4 5
*/
type Pusher interface {
	Push(v any)
}

/*
Popper 弹出

0 1 2 3 4 5→
0 1 2 3 4
*/
type Popper interface {
	Pop() any
}

/*
Puller 拉出

←0 1 2 3 4 5
   1 2 3 4 5
*/
type Puller interface {
	Pull() any
}

/*
Mover 移动
*/
type Mover interface {
	Move(i, j int)
}

/*
Indexer 索引
*/
type Indexer interface {
	Set(i int, v any)
	Get(i int) any
}

package list

import (
	. "github.com/wolfired/golabs/container"
)

type any = interface{}

type node struct {
	nxt  *node
	value any
}

/*
Lister 列表
*/
type Lister interface {
	Sizer
	Indexer
	Inserter
	Deleter
	Pos(v any) int
}

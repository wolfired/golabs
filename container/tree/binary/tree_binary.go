package binary

import (
	"github.com/wolfired/golabs/container/tree"
)

/*
New 新建一棵二叉树
*/
func New() tree.Treer {
	return &binary{}
}

type node struct {
	left  *node
	right *node
	value interface{}
}

type binary struct {
	root *node
}

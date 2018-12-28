package list

import (
	"github.com/wolfired/golabs/container"
)

/*
Lister 列表
*/
type Lister interface {
	container.Sizer
	container.Indexer
	container.Inserter
	container.Deleter
}

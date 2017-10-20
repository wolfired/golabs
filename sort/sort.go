//Package sort 实现各种排序算法
package sort

/*
Comparer 比较
*/
type Comparer interface {
	Compare(i, j uint) int
}

/*
Swapper 交换
*/
type Swapper interface {
	Swap(i, j uint)
}

/*
Sorter 排序
*/
type Sorter interface {
}

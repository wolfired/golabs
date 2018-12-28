package quick

/*
Sort 快速排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	theorySplitV1(sample)
}

func theorySplitV1(unsort []int) {
	n := len(unsort)

	if 2 > n {
		return
	}

	i := 0
	for j := 1; j < n; j++ {
		if unsort[i] > unsort[j] {
			unsort[i], unsort[j] = unsort[j], unsort[i]
			i++
		}
	}

	theorySplitV1(unsort[:i])
	theorySplitV1(unsort[i+1:])
}

func optimize(sample []int) {
	theory(sample)
}

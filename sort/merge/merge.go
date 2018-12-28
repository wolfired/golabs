package merge

/*
Sort 归并排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	unsort := make([]int, len(sample))
	copy(unsort, sample)
	theorySplitV2(unsort, sample)
}

/*
递归实现
*/
func theorySplitV1(unsort []int, sorted []int) {
	n := len(unsort)

	if 2 > n {
		return
	}

	m := n / 2
	theorySplitV1(sorted[:m], unsort[:m])
	theorySplitV1(sorted[m:], unsort[m:])
	theoryMerge(unsort, m, sorted)
}

/*
迭代实现
*/
func theorySplitV2(unsort []int, sorted []int) {
	n := len(unsort)
	s := 1
	for ; s < n; s *= 2 {
		for b := 0; b < n; b = b + 2*s {
			m := b + s
			if n < m {
				m = n
			}
			e := b + 2*s
			if n < e {
				e = n
			}
			theoryMerge(unsort[b:e], m-b, sorted[b:e])
		}
		unsort, sorted = sorted, unsort
	}
	if 0 == s%2 {
		copy(sorted, unsort)
	}
}

func theoryMerge(unsort []int, m int, sorted []int) {
	e := len(unsort)
	for i, j, k := 0, 0, m; i < e; i++ {
		if j < m && (e <= k || unsort[j] <= unsort[k]) {
			sorted[i] = unsort[j]
			j++
		} else {
			sorted[i] = unsort[k]
			k++
		}
	}
}

func optimize(sample []int) {
	theory(sample)
}

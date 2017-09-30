package quick

/*
Sort 快速排序
*/
func Sort(sample []int) {
	theory(sample)
}

func theory(sample []int) {
	sorted := make([]int, len(sample))

	theorySort(sample, sorted)

	copy(sample, sorted)
}

func theorySort(sample, sorted []int) {
	n := len(sample)

	if 1 >= n {
		return
	}

	pivot := sample[n-1]
	h, t := 0, n-1
	for i := 0; i < n && h != t; i++ {
		if sample[i] <= pivot {
			sorted[h] = sample[i]
			h++
		} else if sample[i] > pivot {
			sorted[t] = sample[i]
			t--
		}
	}

	sorted[h] = pivot

	copy(sample, sorted)

	theorySort(sample[:h], sorted[:h])
	theorySort(sample[t:], sorted[t:])
}

func optimize(sample []int) {

}

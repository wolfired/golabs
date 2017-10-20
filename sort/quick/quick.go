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
}

func theorySort(sample, sorted []int) {
	n := len(sample)

	if 1 >= n {
		return
	}

	pivot := sample[n-1]
	h, t := 0, n-1

	for i := 0; i < n; i++ {
		if sample[i] < pivot {
			sorted[h] = sample[i]
			h++
		} else if sample[i] > pivot {
			sorted[t] = sample[i]
			t--
		}
	}

	for i := h; i <= t; i++ {
		sorted[i] = pivot
	}

	copy(sample, sorted)

	theorySort(sample[0:h], sorted[0:h])
	theorySort(sample[t+1:], sorted[t+1:])
}

func optimize(sample []int) {
	n := len(sample)

	if 1 >= n {
		return
	}

	pivot := sample[n-1]
	i, h, t := 0, 0, n-1
	for {
		if sample[i] < pivot {
			h++
			i++
		} else if sample[i] > pivot {
			sample[i], sample[t] = sample[t], sample[i]
			t--
		}
	}
}

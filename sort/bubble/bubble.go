package bubble

/*
Sort 冒泡排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	n := uint(len(sample))

	for i := uint(1); i < n; i++ {
		for j := uint(1); j < n; j++ {
			if sample[j-1] > sample[j] {
				sample[j-1], sample[j] = sample[j], sample[j-1]
			}
		}
	}
}

func optimize(sample []int) {
	n := uint(len(sample))

	for sorted, i := false, uint(1); !sorted && i < n; i++ {
		sorted = true

		for j := uint(1); j < n-i+1; j++ {
			if sample[j-1] > sample[j] {
				sorted = false
				sample[j-1], sample[j] = sample[j], sample[j-1]
			}
		}
	}
}

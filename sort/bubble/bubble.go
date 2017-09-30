package bubble

/*
Sort 冒泡排序
*/
func Sort(sample []int) {
	optimize(sample)
}

func theory(sample []int) {
	n := len(sample)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if sample[j] > sample[j+1] {
				sample[j], sample[j+1] = sample[j+1], sample[j]
			}
		}
	}
}

func optimize(sample []int) {
	n := len(sample)

	for i := 0; i < n-1; i++ {
		sorted := true

		for j := 0; j < n-1-i; j++ {
			if sample[j] > sample[j+1] {
				sorted = false
				sample[j], sample[j+1] = sample[j+1], sample[j]
			}
		}

		if sorted {
			break
		}
	}
}

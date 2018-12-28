package bubble

/*
Sort 冒泡排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	n := len(sample)

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if sample[j-1] > sample[j] {
				sample[j-1], sample[j] = sample[j], sample[j-1]
			}
		}
	}
}

func optimize(sample []int) {
	n := len(sample)

	for unsort, i := true, 1; unsort && i < n; i++ {
		unsort = false

		for j := 1; j < n-i+1; j++ {
			if sample[j-1] > sample[j] {
				unsort = true
				sample[j-1], sample[j] = sample[j], sample[j-1]
			}
		}
	}
}

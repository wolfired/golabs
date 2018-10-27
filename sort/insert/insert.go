package insert

/*
Sort 插入排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
}

func optimize(sample []int) {
	n := len(sample)

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if sample[j-1] > sample[j] {
				sample[j-1], sample[j] = sample[j], sample[j-1]
			} else {
				break
			}
		}
	}
}

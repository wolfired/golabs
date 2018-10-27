package shell

/*
Sort 希尔排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
}

func optimize(sample []int) {
	gaps := [...]int{701, 301, 132, 57, 23, 10, 4, 1}

	n := len(sample)

	for _, gap := range gaps {
		for i := gap; i < n; i++ {
			for j := i; j >= gap; j -= gap {
				if sample[j-gap] > sample[j] {
					sample[j-gap], sample[j] = sample[j], sample[j-gap]
				} else {
					break
				}
			}
		}
	}
}

package shell

/*
Sort 希尔排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	n := len(sample)

	for gap := n / 2; 0 < gap; gap /= 2 {
		for i := gap; i < n; i++ {
			for j := i; gap <= j; j -= gap {
				if sample[j-gap] > sample[j] {
					sample[j-gap], sample[j] = sample[j], sample[j-gap]
				} else {
					break
				}
			}
		}
	}
}

func optimize(sample []int) {
	theory(sample)
}

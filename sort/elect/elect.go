package elect

/*
Sort 选择排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	n := len(sample)

	for i := n - 1; i > 0; i-- {
		for j := i; j > 0; j-- {
			if sample[j-1] > sample[i] {
				sample[j-1], sample[i] = sample[i], sample[j-1]
			}
		}
	}
}

func optimize(sample []int) {
	n := len(sample)

	for i := n - 1; i > 0; i-- {
		c := i

		for j := c; j > 0; j-- {
			if sample[j-1] > sample[c] {
				c = j - 1
			}
		}

		if c != i {
			sample[c], sample[i] = sample[i], sample[c]
		}
	}
}

package elect

/*
Sort 选择排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
}

func optimize(sample []int) {
	n := uint(len(sample))

	for i := uint(1); i < n; i++ {

		c, m := n-i, n-i

		for j := c; j > 0; j-- {
			if sample[j-1] > sample[c] {
				c = j - 1
			}
		}

		if c != m {
			sample[c], sample[m] = sample[m], sample[c]
		}
	}
}

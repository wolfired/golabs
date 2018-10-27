package elect

/*
Sort 选择排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
}

func optimize(sample []int) {
	for i := len(sample) - 1; i > 0; i-- {

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

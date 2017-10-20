package cocktail

/*
Sort 鸡尾酒排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
}

func optimize(sample []int) {
	for h, t := uint(0), uint(len(sample)); h < t; {
		for i := t - 1; i > h; i-- {
			if sample[i-1] > sample[i] {
				sample[i-1], sample[i] = sample[i], sample[i-1]
			}
		}
		h++

		for i := h + 1; i < t; i++ {
			if sample[i-1] > sample[i] {
				sample[i-1], sample[i] = sample[i], sample[i-1]
			}
		}
		t--
	}
}

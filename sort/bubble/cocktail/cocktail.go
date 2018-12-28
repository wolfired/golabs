package cocktail

/*
Sort 鸡尾酒排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	n := len(sample)

	for h, t := 0, n; h < t; {
		for i := t - 1; i > h; i-- {
			if sample[i-1] > sample[i] {
				sample[i-1], sample[i] = sample[i], sample[i-1]
			}
		}
		h++

		for s := h + 1; s < t; s++ {
			if sample[s-1] > sample[s] {
				sample[s-1], sample[s] = sample[s], sample[s-1]
			}
		}
		t--
	}
}

func optimize(sample []int) {
	n := len(sample)

	for unsort, h, t := true, 0, n; unsort && h < t; {
		unsort = false

		for i := t - 1; i > h; i-- {
			if sample[i-1] > sample[i] {
				unsort = true
				sample[i-1], sample[i] = sample[i], sample[i-1]
			}
		}
		h++

		for s := h + 1; unsort && s < t; s++ {
			if sample[s-1] > sample[s] {
				unsort = true
				sample[s-1], sample[s] = sample[s], sample[s-1]
			}
		}
		t--
	}
}

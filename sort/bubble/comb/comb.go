package comb

/*
Sort 梳排序
*/
func Sort(sample []int) {
}

func theory(sample []int) {
	n := len(sample)
	gap, shrink := n, 1.3

	for 1 < n {
		gap = int(float64(gap) / shrink)

		if 1 > gap {
			gap = 1
		}

		for i := 0; i+gap < n; i++ {
			if sample[i] > sample[i+gap] {
				sample[i], sample[i+gap] = sample[i+gap], sample[i]
			}
		}

		if 1 == gap {
			break
		}
	}
}

func optimize(sample []int) {
	theory(sample)
}

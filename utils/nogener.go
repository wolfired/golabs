package utils

import (
	"fmt"
)

func Nogener(kses []uint, h, l uint) (noes []string) {
	noes = make([]string, 0)

	var i int
	var sum uint
	for i = 1; i < len(kses)-2; i++ {
		sum += kses[i]
		sum %= 10

		if sum > h {
			noes = append(noes, nogener(h, l, kses[i-1:i+3]))
		} else if sum > l {
			noes = append(noes, nogener(l, h, kses[i-1:i+3]))
		}
	}

	return
}

func nogener(k uint, d uint, ks []uint) string {
	k0, k1, k2, k3 := ks[0], ks[1], ks[2], ks[3]

	no := [...]uint{(k + k0 + d) % 10, k, (k0 + k1 + d) % 10, (k1 + k2 + d) % 10, (k0 + k2 + d) % 10, k3, (k2 + k3 + d) % 10}

	return fmt.Sprint(no)
}

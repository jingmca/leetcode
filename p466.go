package main

import (
	"fmt"
)

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {

	if len(s2)*n2 > len(s1)*n1 {
		return 0
	}

	pattern := make([]int, n1)
	counter := make([]int, n1)

	i, count := 0, 0

	for block := 0; block < n1; block++ {
		for _, c1 := range s1 {
			if byte(c1) == byte(s2[i]) {
				i++
			}
			if i == len(s2) {
				i = 0
				count++
			}
		}

		pattern[block] = i
		counter[block] = count

		for prev := 0; prev < block; prev++ {

			if pattern[prev] == i && counter[prev] > 0 {
				prev_count := counter[prev]
				repeat_count := (counter[block] - counter[prev]) * ((n1 - 1 - prev) / (block - prev))
				remain_count := counter[prev+(n1-1-prev)%(block-prev)] - counter[prev]

				fmt.Println(pattern, counter, prev, block)

				return (prev_count + repeat_count + remain_count) / n2
			}
		}
	}
	return counter[n1-1] / n2

}

func main() {
	// fmt.Println(getMaxRepetitions("niconiconi", 99981, "nico", 81))
	fmt.Println(getMaxRepetitions("aaa", 3, "aa", 1))
}

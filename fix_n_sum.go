package main

import "fmt"

func sliceSumWithN(s []int, num int, target int) {
	digits := 1 << uint(len(s))

	for i := 1; i < digits; i++ {
		sum := 0
		ele := make([]int, 0, num)

		if n1of(i) == num {
			for j := 0; j < len(s); j++ {
				if (i & (1 << uint(j))) != 0 {
					sum += s[j]
					ele = append(ele, s[j])
				}
			}

			if sum == target {
				fmt.Printf("found :%v\n", ele)
			}
		}
	}
}

func n1of(x int) int {
	count := 0
	for x > 0 {
		x = x & (x - 1)
		count++
	}

	return count
}

func main() {
	sliceSumWithN([]int{2, 3, 4, 5, 1}, 3, 10)
}

package main

import (
	"fmt"
)

func stoneGame(piles []int) bool {

	result := make([]int, len(piles))
	for i, val := range piles {
		result[i] = val
	}

	for j := 1; j < len(piles); j++ {
		for i := 0; i+j < len(piles); i++ {
			result[i+j] = max(piles[i]-result[i+j], piles[i+j]-result[i+j-1])
		}
	}

	return result[0] > 0
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(stoneGame([]int{1, 4, 10, 8, 3, 2, 4, 1}))
}

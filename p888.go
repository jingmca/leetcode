package main

import (
	"fmt"
)

func fairCandySwap(A []int, B []int) []int {

	candys := append(A, B...)
	sum := 0
	for _, val := range candys {
		sum += val
	}
	size := sum / 2
	sumA := 0
	for _, val := range A {
		sumA += val
	}
	mapB := make(map[int]int)
	for j, val := range B {
		mapB[val] = j
	}

	for _, val := range A {
		if _, ok := mapB[val+size-sumA]; ok {
			return []int{val, val + size - sumA}
		}
	}
	return []int{}
}

func main() {
	fmt.Println(fairCandySwap([]int{20, 35, 22, 6, 13}, []int{31, 57}))
}

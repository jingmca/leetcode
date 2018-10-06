package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	minCost := make([]int, 0, len(cost)+1)
	minCost = append(minCost, 0, 0)

	for i := 2; i <= len(cost); i++ {
		if minCost[i-2]+cost[i-2] < minCost[i-1]+cost[i-1] {
			minCost = append(minCost, minCost[i-2]+cost[i-2])
		} else {
			minCost = append(minCost, minCost[i-1]+cost[i-1])
		}
	}

	fmt.Println(minCost)
	return minCost[len(minCost)-1]
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}

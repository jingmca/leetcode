package main

import (
	"fmt"
)

func climbStairs(n int) int {

	climbs := make([]int, 0, n+1)
	climbs = append(climbs, 0, 1, 2)

	for i := 3; i <= n; i++ {
		climbs = append(climbs, climbs[i-1]+climbs[i-2])
	}
	return climbs[n]

}

func main() {
	fmt.Println(climbStairs(4))
}

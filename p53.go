package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sum := nums[0]
	max := nums[0]

	for _, i := range nums[1:] {
		if sum+i > i {
			sum += i
		} else {
			sum = i
		}
		if sum > max {
			max = sum
		}
	}

	return max

}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1}))
}

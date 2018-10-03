package main

import "fmt"

func maxProfit(prices []int) int {

	profit := 0

	for i, left := range prices {
		for _, right := range prices[i:] {
			if right-left > profit {
				profit = right - left
			}
		}
	}

	return profit

}

func maxProfit2(prices []int) int {

	profit := 0
	min := prices[0]

	for _, left := range prices {
		if left < min {
			min = left
		} else if left-min > profit {
			profit = left - min
		}
	}

	return profit

}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x >= 0 && x <= 9 {
		return true
	} else {
		y := 0
		x1 := x
		for x > 0 {
			digit := x % 10
			x /= 10
			y = 10*y + digit
		}
		return y == x1
	}
}

func main() {
	fmt.Println(isPalindrome(121))
}

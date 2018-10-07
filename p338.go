package main

import (
	"fmt"
)

func countBits(num int) []int {

	bits := make([]int, num+1)
	bits[0] = 0
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + (i % 2)
	}

	return bits

}

func main() {
	fmt.Println(countBits(8))
}

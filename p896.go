package main

func isMonotonic(A []int) bool {
	if len(A) < 2 {
		return true
	}

	ret := 1
	if A[len(A)-1] <= A[0] {
		ret = 0
	}

	for i := 1; i < len(A); i++ {
		switch ret {
		case 1:
			if A[i] < A[i-1] {
				return false
			}
		case 0:
			if A[i] > A[i-1] {
				return false
			}
		}
	}

	return true
}

// func main() {
// 	fmt.Println(isMonotonic([]int{4, 4, 4, 4}))
// }

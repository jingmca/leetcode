package main

import (
	"fmt"
	"sort"
	"strings"
)

func orderlyQueue(S string, K int) string {
	l := len(S)
	if l <= 1 || K == 0 {
		return S
	}

	if K >= l {
		K = l
	}

	if K == 1 {
		s := S
		for i := 1; i < l; i++ {
			t := ""
			if i == l-1 {
				t = fmt.Sprintf("%s%s", string(S[i]), S[:i])
			} else {
				t = fmt.Sprintf("%s%s", S[i:], S[:i])
			}

			if strings.Compare(s, t) > 0 {
				s = t
			}
		}

		S = s

	} else {
		s := strings.Split(S, "")
		sort.Strings(s)
		S = strings.Join(s, "")
	}

	return S

}

func minIndex(s string, excpte int) int {
	min := s[0]
	index := 0
	for i, c := range s {
		if byte(c) < min {
			if excpte == -1 || i != excpte {
				min = byte(c)
				index = i
			}
		}
	}

	return index
}

// func main() {

// 	fmt.Println(orderlyQueue("xxqjzq", 2))
// }

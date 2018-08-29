package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}

	rstr := ""
	cstr := fmt.Sprintf("%c", s[0])
	for _, i := range s[1:] {
		k := strings.LastIndex(cstr, string(i))
		if k == -1 {
			cstr += fmt.Sprintf("%s", string(i))
		} else {
			if len(cstr) > len(rstr) {
				rstr = cstr
			}
			cstr = cstr[k+1:] + string(i)
		}

		fmt.Println(rstr, cstr)
	}

	fmt.Println(rstr, cstr)

	if len(rstr) >= len(cstr) {
		return len(rstr)
	} else {
		return len(cstr)
	}
}

// func main() {
// 	a := lengthOfLongestSubstring("dababavdf")
// 	fmt.Println(a)
// }

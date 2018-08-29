package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	if len(s) == 2 {
		if s[0] == s[1] {
			return s
		} else {
			return string(s[0])
		}
	}

	tailLen := 0
	maxLen := 0
	tailStr := ""
	maxStr := ""
	start := 0

	for i := 0; i < len(s); i++ {
		if i == 0 {
			continue
		}

		if maxLen <= 3 {
			start = i - 4
		} else {
			start = i - maxLen - 1
		}

		if start < 0 {
			start = 0
		}

		tailLen, tailStr = tailPalindrome(s[start : i+1])
		if tailLen >= maxLen {
			maxLen = tailLen
			maxStr = tailStr
		}

		fmt.Printf("step %d, maxlen=%d, maxstr=%q, tailen=%d, tailstr=%q\n", i, maxLen, maxStr, tailLen, tailStr)

	}

	if maxLen == 0 {
		return string(s[0])
	} else {
		return maxStr
	}

}

func tailPalindrome(s string) (slen int, str string) {
	if len(s) < 2 {
		return 0, ""
	}

	slen = len(s)
	str = s
	stop := 0
	if slen%2 == 0 {
		stop = (slen - 2) / 2
	} else {
		stop = (slen - 1) / 2
	}

	ret := true
	for i := 0; i <= stop; i++ {
		if s[i] != s[slen-i-1] {
			ret = false
			break
		}
	}

	if ret {
		return slen, str
	} else {
		return tailPalindrome(s[1:])
	}
}

func main() {
	s := longestPalindrome("a")
	fmt.Println(s)
}

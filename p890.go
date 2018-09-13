package main

import "fmt"

func findAndReplacePattern(words []string, pattern string) []string {

	ret := make([]string, 0, len(words))

	for _, w := range words {
		if len(w) == len(pattern) {
			wcmap := make(map[rune]byte)
			cwmap := make(map[byte]rune)
			flag := true
			for i, c := range w {
				val, ok := wcmap[c]
				if ok {
					if val != pattern[i] {
						flag = false
						break
					} else {
						if val2, ok2 := cwmap[val]; ok2 {
							if val2 != c {
								flag = false
								break
							}
						} else {
							flag = false
							break
						}
					}
				} else {
					wcmap[c] = pattern[i]
					if val2, ok2 := cwmap[pattern[i]]; ok2 {
						if val2 != c {
							flag = false
							break
						}
					} else {
						cwmap[pattern[i]] = c
					}
				}
			}
			if flag {
				ret = append(ret, w)
			}
		}
	}

	return ret

}

func main() {
	fmt.Println(findAndReplacePattern([]string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}, "abb"))
}

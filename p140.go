package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {

	if len(s) == 0 {
		return []string{}
	}

	alphabets := make(map[rune]bool)
	for _, w := range wordDict {
		for _, c := range w {
			alphabets[c] = true
		}
	}

	for _, c := range s {
		if _, ok := alphabets[c]; !ok {
			return []string{}
		}
	}

	var record = make(map[string][]string)
	return wordBreak2(s, wordDict, &record)
}

func wordBreak2(s string, wordDict []string, rDict *map[string][]string) []string {
	res := make([]string, 0)

	for _, w := range wordDict {
		head := strings.Index(s, w)

		if head == 0 {
			_s := s[len(w):]
			if len(_s) == 0 {
				res = append(res, w)
			} else {
				if _r, ok := (*rDict)[_s]; !ok {
					_r = wordBreak2(_s, wordDict, rDict)
					for _, rs := range _r {
						res = append(res, fmt.Sprintf("%s %s", w, rs))
					}

				} else {
					for _, rs := range _r {
						res = append(res, fmt.Sprintf("%s %s", w, rs))
					}
				}

			}

		}
	}
	(*rDict)[s] = res
	return res
}

func main() {
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))

	// fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba"}))
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

}

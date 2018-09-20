package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {

	endpoints := make(map[int][]int)

	for _, w := range wordDict {

		for i, width := 0, len(w); strings.Contains(s[i:], w); {
			idx := strings.Index(s[i:], w)
			if idx > -1 {
				endpoints[idx+i] = append(endpoints[idx+i], idx+i+width)
				i += idx + 1
			}
		}
	}

	stop := len(s)

	del := make(map[int][]int)
	for k, v := range endpoints {
		for _, iv := range v {
			if _, ok := endpoints[iv]; !ok && iv != stop {
				del[k] = append(del[k], iv)
			}
		}
	}

	fmt.Println(del)

	for k, v := range del {
		for _, iv := range v {
			for idx, iiv := range endpoints[k] {
				if iiv == iv {
					endpoints[k] = append(endpoints[k][:idx], endpoints[k][idx+1:]...)
				}
			}
		}
	}

	fmt.Println(endpoints)

	res := make([]string, 0)

	calc(0, stop, []int{}, endpoints, &res, s)

	// fmt.Println(output)

	// for _, o := range output {
	// 	_str := make([]string, 0, len(o)-1)
	// 	for i := 0; i < len(o)-1; i++ {
	// 		_str = append(_str, s[o[i]:o[i+1]])
	// 	}
	// 	fmt.Println(strings.Join(_str, " "), o)
	// 	res = append(res, strings.Join(_str, " "))
	// }

	return res
}

func calc(start int, stop int, prev []int, ends map[int][]int, res *[]string, s string) {
	if next, ok := ends[start]; ok {
		for _, idx := range next {
			if idx == stop {
				_out := append(prev, start, idx)
				// *out = append(*out, _out)
				// fmt.Printf("%v, %d, %d, %v\n", _out, idx, stop, out)
				_str := make([]string, 0, len(_out)-1)
				for i := 0; i < len(_out)-1; i++ {
					_str = append(_str, s[_out[i]:_out[i+1]])
				}
				// fmt.Println(strings.Join(_str, " "), _out)
				*res = append(*res, strings.Join(_str, " "))

			} else {
				path := append(prev, start)
				calc(idx, stop, path, ends, res, s)
			}
		}
	}
}

func main() {
	fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))
}

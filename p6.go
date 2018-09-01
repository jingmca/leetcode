package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows < 2 {
		return s
	}

	fmt.Println(len(s))

	i := len(s)
	last := 0
	strList := make([]string, 0, i)

	for j := 0; i > 0; j++ {
		if last < numRows {
			if i >= numRows {
				strList = append(strList, s[len(s)-i:len(s)-i+numRows])
				// strList[j] = s[len(s)-i : len(s)-i+numRows]
				i -= numRows
				last = numRows
			} else {
				strList = append(strList, s[len(s)-i:])
				// strList[j] = s[len(s)-i:]
				i = 0
				last = i
			}
		} else {
			if i >= numRows-2 {
				extend := k2(s[len(s)-i:len(s)-i+numRows-2], numRows)
				for _, k := range extend {
					strList = append(strList, k)
				}
				i -= numRows - 2
				last = numRows - 2
			} else {
				extend := k2(s[len(s)-i:], numRows)
				for _, k := range extend {
					strList = append(strList, k)
				}
				i = 0
				last = i
			}
		}

	}

	str := ""
	for i := 0; i < numRows; i++ {
		for _, c := range strList {
			if i <= len(c)-1 && string(c[i]) != " " {
				str += string(c[i])
			}
		}
	}
	return str
}

func k2(s string, row int) []string {
	sList := make([]string, len(s))

	for i, c := range s {
		ts := ""
		for j := 0; j < row-(i+2); j++ {
			ts += " "
		}
		ts += string(c)
		for j := 0; j < i+1; j++ {
			ts += " "
		}
		sList[i] = ts
	}
	return sList
}

// func main() {
// 	s := convert("yjvsbxetkierlqfbxyetjbyqqgrtrurpfmkhjocwyjpjzunxsrqdurtkxngqjxgokqxvgarejgqkadhuuayortbqgjhpgpgsfdolffrqmhbokklgklxdeywnhfepukibqwoxbfqpnrgzdrgocdtidpxmucbqojrghfelnuaangzszhibmcmptjbqnfgcoykyfojskluzuwstdaxqejhyuloiqumguwecnnuzbpzvntoqvliawsatdobtkpzhlejytfauwzrjugcptmrserlhhoaudcboimpdrpaqqrzmxddtqvluoweymgspitfshwwynwqfnqrnvvilstiirmgduyuftzxawvbjrrphjiwffgszzcisqoxlprqkqnloloaehrltzjahpsgqxoknfhywyogrethphhtrahkcsmjkgpcdgnrnwpjxgpqkjxbshwlhfxjyjskqkmtqbkdycovidnuokvjrtubzukzdfjtpxphzzmzbawlfjfuvcfpwbqxvcyzhhuygjhhltgoteaznhvlkaaidqhzsfacoucwekerfmfzrhagpxrbxtlajsbezbgnwklcupvaeabviddxaxazqlbcddgcgoreacixudzyeavofanfxngqyhubmaftqyzqcinylowrotfvusctfijdsdggfnbxnbqsjfqwupokitgcmiwtthxlnfruvtsiuiafprbjgpuq", 415)
// 	fmt.Println(s)
// }

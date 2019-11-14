package leetcode

import "strings"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	strs := strings.Split(digits, "")
	ret := m[strs[0]]
	for i := 1; i < len(strs); i++ {
		temp := m[strs[i]]
		stemp := []string{}
		for j := 0; j < len(ret); j++ {
			for _, s := range temp {
				stemp = append(stemp, ret[j]+s)
			}
		}
		ret = stemp
	}
	return ret
}

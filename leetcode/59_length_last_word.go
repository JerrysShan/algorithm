package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	strs := strings.Split(s, " ")
	return len(strs[len(strs)-1])
}

package leetcode

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	reg, err := regexp.Compile(`[\W\s]*`)
	if err != nil {
		panic(err)
	}
	t := strings.ToLower(reg.ReplaceAllString(s, ""))
	length := len(t)
	for i := 0; i < length/2; i++ {
		if t[i] != t[length-1-i] {
			return false
		}
	}
	return true
}

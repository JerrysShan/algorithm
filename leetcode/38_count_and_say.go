package leetcode

import "strconv"

func countAndSay(n int) string {
	str := "1"
	if n == 1 {
		return str
	}
	for i := 2; i <= n; i++ {
		tmp := ""
		for j := 0; j < len(str); j++ {
			count := 1
			for j+1 < len(str) && str[j+1] == str[j] {
				count++
				j++
			}
			tmp += strconv.Itoa(count) + string(str[j])
		}
		str = tmp
	}
	return str
}

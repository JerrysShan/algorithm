package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	i := 0
	j := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			ret++
			i++
		}
		j++
	}
	return ret
}

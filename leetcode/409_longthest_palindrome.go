package leetcode

func longestPalindrome2(s string) int {
	m := make(map[rune]int)
	for _, v := range s {
		if count, ok := m[v]; ok {
			m[v] = count + 1
		} else {
			m[v] = 1
		}
	}
	ret := 0
	for k, v := range m {
		if v%2 == 0 {
			ret += v
			m[k] = 0
		} else {
			ret += v - 1
			m[k] = 1
		}
	}
	for _, v := range m {
		if v == 1 {
			ret += v
			return ret
		}
	}
	return ret
}

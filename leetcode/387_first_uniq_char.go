package leetcode

func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for index, v := range s {
		if _, ok := m[v]; ok {
			m[v] = -1
		} else {
			m[v] = index
		}
	}
	first := -1
	for _, v := range m {
		if v != -1 && (first == -1 || v < first) {
			first = v
		}
	}
	return first
}

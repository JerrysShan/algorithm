package leetcode

func isValid(s string) bool {
	m := map[string]string{"}": "{", "]": "[", ")": "("}
	for index := range s {
		if v, ok := m[string(s[index])]; ok {
			if index == 0 || string(s[index-1]) != v {
				return false
			}
		}
	}
	return true
}

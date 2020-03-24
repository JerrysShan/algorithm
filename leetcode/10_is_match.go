package leetcode

func isMatch(s string, p string) bool {
	if p == "" {
		if s == "" {
			return true
		}
		return false
	}
	firstMatch := s != "" && (p[0] == s[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	}
	return firstMatch && isMatch(s[1:], p[1:])
}

package leetcode

func isPalindromic(s string) bool {
	length := len(s)
	for k := 0; k < length/2; k++ {
		if s[k] != s[length-k-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	lstr := ""
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			test := s[i:j]
			if isPalindromic(test) && len(test) > max {
				lstr = test
				max = len(test)
			}
		}
	}
	return lstr
}

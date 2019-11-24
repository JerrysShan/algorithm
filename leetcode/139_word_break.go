package leetcode

func wordBreak(s string, wordDict []string) bool {
	return wordBreakFun(s, wordDict, 0, make([]bool, len(s)))
}

func wordBreakFun(s string, wordDict []string, start int, mem []bool) bool {
	if start == len(s) {
		return true
	}
	for end := start + 1; end <= len(s); end++ {
		t := s[start:end]
		if mem[start] {
			return true
		}
		isExist := false
		for _, v := range wordDict {
			if v == t {
				isExist = true
			}
		}
		if isExist && wordBreakFun(s, wordDict, end, mem) {
			mem[start] = true
			return true
		}
	}
	mem[start] = false
	return false
}

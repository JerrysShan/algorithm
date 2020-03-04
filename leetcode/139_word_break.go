package leetcode

func wordBreak(s string, wordDict []string) bool {
	return wordBreakFun(s, wordDict, 0, make([]bool, len(s)))
}

func wordBreakFun(s string, wordDict []string, start int, mem []bool) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			isExist := false
			for _, v := range wordDict {
				if v == s[j:i] {
					isExist = true
				}
			}
			if dp[j] && isExist {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func wordBreakFun2(s string, wordDict []string, start int) bool {
	if start == len(s) {
		return true
	}
	for end := start + 1; end <= len(s); end++ {
		isContain := false
		for _, v := range wordDict {
			if v == s[start:end] {
				isContain = true
				break
			}
		}
		if isContain && wordBreakFun2(s, wordDict, end) {
			return true
		}
	}
	return false
}

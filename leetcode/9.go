package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n := 0
	t := x
	for t != 0 {
		n = n*10 + t%10
		t = t / 10
	}
	if n == x {
		return true
	}
	return false
}

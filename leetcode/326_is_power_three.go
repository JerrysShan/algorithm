package leetcode

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	for n != 0 {
		if n%3 != 0 {
			return false
		}
		if n/3 == 1 && n%3 == 0 {
			return true
		}
		n = n / 3
	}
	return false
}

package leetcode

func isPowerOfTwo(num int) bool {
	for num >= 2 {
		if num%2 == 0 {
			num = num / 2
		} else {
			return false
		}

	}
	if num == 1 {
		return true
	}
	return false
}

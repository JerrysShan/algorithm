package leetcode

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	var ret float64 = 1
	for i := 0; i < n; i++ {
		ret *= x
	}
	return ret
}

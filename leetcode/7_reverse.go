package leetcode

import "go-go1.6/src/math"

//整数的反转
func reverse(x int) int {
	ret := 0
	for x != 0 {
		c := x % 10
		ret = ret*10 + c
		if ret > math.MaxInt32 || ret < math.MinInt32 {
			return 0
		}
		x = x / 10
	}

	return ret
}

package leetcode

import "math"

func divide(dividend int, divisor int) int {
	result := 0
	sign := (dividend > 0) == (divisor > 0)
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	count := 0
	for dividend >= divisor {
		count++
		divisor = divisor << 1
	}
	for count > 0 {
		count--
		divisor = divisor >> 1
		if divisor <= dividend {
			result += 1 << uint(count)
			dividend -= divisor
		}
	}
	if !sign {
		result = -result
	}
	if result < -(1<<31) || result > 1<<31-1 {
		return 1<<31 - 1
	}
	return result
}

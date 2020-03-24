package leetcode

func findNthDigit(n int) int {
	base := 9
	digits := 1
	for n-base*digits > 0 {
		n = n - base*digits
		digits++
		base = base * 10
	}
	idx := n % digits
	if idx == 0 {
		idx = digits
	}
	number := 1
	for i := 1; i < digits; i++ {
		number *= 10
	}
	if idx == digits {
		number += n/digits - 1
	} else {
		number += n / digits
	}
	for i := idx; i < digits; i++ {
		number /= 10
	}
	return number % 10
}

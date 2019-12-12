package leetcode

func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return nil
	}
	temp := 0
	if digits[length-1]+1 == 10 {
		temp = 1
		digits[length-1] = 0
	} else {
		digits[length-1] = digits[length-1] + 1
		return digits
	}
	if length == 1 {
		digits = append([]int{1}, digits...)
		return digits
	}
	for i := length - 2; i > -1; i-- {
		if digits[i]+temp == 10 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i] = digits[i] + temp
			break
		}
	}
	return digits
}

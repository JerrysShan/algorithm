package leetcode

func missingNumber(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	max := nums[0]
	m := make(map[int]bool)
	m[max] = true
	var ret int = -1
	for i := 1; i < length; i++ {
		if nums[i] > max {
			max = nums[i]
		}
		m[nums[i]] = true
	}
	for j := 0; j <= max+1; j++ {
		if _, ok := m[j]; !ok {
			ret = j
			break
		}
	}
	return ret
}

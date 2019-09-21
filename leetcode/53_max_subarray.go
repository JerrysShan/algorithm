package leetcode

func maxSubArray(nums []int) int {
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		if temp > result {
			result = temp
		}
		for j := i + 1; j < len(nums); j++ {
			temp += nums[j]
			if temp > result {
				result = temp
			}
		}
	}
	return result
}

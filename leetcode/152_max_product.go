package leetcode

// 最大乘积子序列
func maxProduct(nums []int) int {
	max := nums[0]
	length := len(nums)
	for i := 0; i < length; i++ {
		temp := nums[i]
		if temp > max {
			max = temp
		}
		for j := i + 1; j < length; j++ {
			temp *= nums[j]
			if temp > max {
				max = temp
			}
		}
	}
	return max
}

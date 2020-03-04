package leetcode

func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	k := 1
	for index, v := range nums {
		ret[index] = k
		k = k * v
	}
	k = 1
	for i := len(nums) - 1; i >= 0; i-- {
		ret[i] = ret[i] * k
		k = k * nums[i]
	}
	return ret
}

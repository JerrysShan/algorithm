package leetcode

func permute(nums []int) [][]int {
	result := [][]int{}
	backtrack(len(nums), nums, &result, 0)
	return result
}

func backtrack(n int, nums []int, result *[][]int, first int) {
	if first == n {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
	}
	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(n, nums, result, first+1)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

package leetcode

import "sort"

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	length := len(nums)
	return nums[length-k]
}

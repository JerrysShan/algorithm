package leetcode

func searchInsert(nums []int, target int) int {
	for index, v := range nums {
		if v >= target {
			return index
		}
	}
	return len(nums)
}

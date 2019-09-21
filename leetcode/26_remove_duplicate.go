package leetcode

func removeDuplicates(nums []int) int {
	length := 0
	for i := 0; i < len(nums); i++ {
		isEqual := false
		for j := 0; j <= length; j++ {
			if nums[j] == nums[i] {
				isEqual = true
			}
		}
		if !isEqual {
			length++
			nums[length] = nums[i]
		}

	}
	return length + 1
}

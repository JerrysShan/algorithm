package leetcode

func rotate(nums []int, k int) {
	length := len(nums)
	for i := 0; i < k; i++ {
		last := nums[length-1]
		for j := length - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

package leetcode

func rob(nums []int) int {
	prevMax := 0
	curMax := 0
	for _, v := range nums {
		temp := curMax
		if prevMax+v > curMax {
			curMax = prevMax + v
		}
		prevMax = temp
	}
	return curMax
}

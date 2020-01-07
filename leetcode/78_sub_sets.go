package leetcode

func subsets(nums []int) [][]int {
	ret := [][]int{make([]int, 0)}
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 1; j <= length-i; j++ {
			arr := []int{}
			for k := 0; k < j && i+k < length; k++ {
				arr = append(arr, nums[i+k])
			}
			ret = append(ret, arr)
		}
	}
	return ret
}

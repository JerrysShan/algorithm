package leetcode

func majorityElement(nums []int) int {
	ret := nums[0]
	major := len(nums) / 2
	container := make(map[int]int)
	for _, v := range nums {
		if count, ok := container[v]; ok {
			container[v] = count + 1
		} else {
			container[v] = 1
		}
	}
	for k, v := range container {
		if v > major {
			ret = k
		}
	}
	return ret
}

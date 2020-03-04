package leetcode

func containsDuplicate(nums []int) bool {
	ret := false
	container := make(map[int]int)
	for _, v := range nums {
		if _, ok := container[v]; ok {
			ret = true
			break
		} else {
			container[v] = 1
		}
	}
	return ret
}

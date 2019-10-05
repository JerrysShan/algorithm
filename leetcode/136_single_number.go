package leetcode

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}

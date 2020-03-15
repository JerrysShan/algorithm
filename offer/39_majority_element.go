package offer

func majorityElement(nums []int) int {
	m := make(map[int]int)
	length := len(nums)

	for _, v := range nums {
		if total, ok := m[v]; ok {
			if (total + 1) > length/2 {
				return v
			}
			m[v] = total + 1
		} else {
			m[v] = 1
		}
	}
	return 0
}

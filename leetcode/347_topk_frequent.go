package leetcode

func topKFrequent(nums []int, k int) []int {
	ret := make([]int, k)
	m := make(map[int]int)
	for _, v := range nums {
		if count, ok := m[v]; ok {
			m[v] = count + 1
		} else {
			m[v] = 1
		}
	}
	i := 0
	for key, val := range m {
		if i < k {
			ret[i] = key
			i++
		} else {
			selectSort(ret, m)
			if val > m[ret[0]] {
				ret[0] = key
			}
		}
	}
	return ret
}

func selectSort(arr []int, m map[int]int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if m[arr[j]] < m[arr[min]] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

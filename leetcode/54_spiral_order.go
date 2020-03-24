package leetcode

func spiralOrder(matrix [][]int) []int {
	ret := []int{}
	n := len(matrix)
	if n == 0 {
		return ret
	}
	m := len(matrix[0])
	t := 0
	r := m - 1
	b := n - 1
	l := 0
	for l <= r && t <= b {
		for i := l; i <= r; i++ {
			ret = append(ret, matrix[t][i])
		}
		for j := t + 1; j <= b; j++ {
			ret = append(ret, matrix[j][r])
		}
		if l < r && t < b {
			for i := r - 1; i > l; i-- {
				ret = append(ret, matrix[b][i])
			}
			for j := b; b > t; j-- {
				ret = append(ret, matrix[j][l])
			}
		}
		t++
		r--
		b--
		l++
	}
	return ret
}

package leetcode

func uniquePaths(m int, n int) int {
	dirs := make([][]int, m)
	for i := 0; i < len(dirs); i++ {
		dirs[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dirs[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dirs[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dirs[i][j] = dirs[i-1][j] + dirs[i][j-1]
		}
	}
	return dirs[m-1][n-1]
}

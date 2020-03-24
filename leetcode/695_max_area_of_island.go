package leetcode

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := getArea(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func getArea(grid [][]int, i, j int) int {
	if i == len(grid) || i < 0 {
		return 0
	}
	if j == len(grid[0]) || j < 0 {
		return 0
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
		return 1 + getArea(grid, i+1, j) + getArea(grid, i-1, j) + getArea(grid, i, j+1) + getArea(grid, i, j-1)
	}
	return 0
}

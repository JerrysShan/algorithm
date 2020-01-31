package leetcode

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := make([][]int, 3)
	grid[0] = []int{1, 3, 1}
	grid[1] = []int{1, 5, 1}
	grid[2] = []int{4, 2, 1}
	t.Log(minPathSum(grid))
}

package leetcode

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	t.Log(exist(board, "ABCCED"))
	t.Log(exist(board, "SEE"))
	t.Log(exist(board, "ABCB"))
}

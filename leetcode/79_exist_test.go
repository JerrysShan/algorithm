package leetcode

import "testing"

func TestExist(t *testing.T) {
	board := make([][]byte, 3)
	board[0] = []byte{'A', 'B', 'C', 'E'}
	board[1] = []byte{'S', 'F', 'C', 'S'}
	board[2] = []byte{'A', 'D', 'E', 'E'}
	t.Log(exist(board, "ABCCED"))
	t.Log(exist(board, "SEE"))
	t.Log(exist(board, "ABCB"))
}

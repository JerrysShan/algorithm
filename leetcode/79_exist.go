package leetcode

func exist(board [][]byte, word string) bool {
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	marked := make([][]bool, m)
	for i := 0; i < m; i++ {
		marked[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if searchWord(board, word, 0, i, j, m, n, marked) {
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte, word string, index, startX, startY, m, n int, marked [][]bool) bool {
	if index == len(word)-1 {
		return board[startX][startY] == word[index]
	}
	directions := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	if board[startX][startY] == word[index] {
		marked[startX][startY] = true
		for _, v := range directions {
			newX := startX + v[0]
			newY := startY + v[1]
			if newX >= 0 && newX < m && newY >= 0 && newY < n && !marked[newX][newY] {
				if searchWord(board, word, index+1, newX, newY, m, n, marked) {
					return true
				}
			}
		}
		marked[startX][startY] = false
	}

	return false
}

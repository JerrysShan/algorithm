package leetcode

func exist(board [][]byte, word string) bool {
	if board == nil || word == "" {
		return false
	}
	m := make(map[byte][]byte)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			char := board[i][j]
			m[char] = []byte{}
			if i-1 > -1 {
				m[char] = append(m[char], board[i-1][j])
			}
			if i+1 < len(board) {
				m[char] = append(m[char], board[i+1][j])
			}
			if j-1 > -1 {
				m[char] = append(m[char], board[i][j-1])
			}
			if j+1 < len(board[i]) {
				m[char] = append(m[char], board[i][j+1])
			}
		}
	}

	bufs := []byte(word)
	for index := 0; index < len(bufs); index++ {
		c := bufs[index]
		if val, ok := m[c]; !ok {
			return false
		} else if index+1 < len(bufs) {
			if !existsChar(bufs[index+1], val) {
				return false
			}
		}
	}
	return true
}

func existsChar(char byte, bytes []byte) bool {
	if bytes == nil {
		return false
	}
	for _, v := range bytes {
		if v == char {
			return true
		}
	}
	return false
}

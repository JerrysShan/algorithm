package offer

import (
	"math"
)

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	m := make(map[rune]int)
	for index, char := range s {
		if _, ok := m[char]; ok {
			m[char] = -1
		} else {
			m[char] = index
		}
	}
	index := math.MaxInt32
	ret := byte(' ')
	for k, v := range m {
		if v != -1 && v < index {
			index = v
			ret = byte(k)
		}
	}
	return ret
}

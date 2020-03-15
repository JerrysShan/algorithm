package offer

import (
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return false
	}
	return true
}

package offer

import "testing"

func TestIsNumber(t *testing.T) {
	t.Log(isNumber("5e2"))
	t.Log(isNumber("12e"))
	t.Log(isNumber("-1E-16"))
	t.Log(isNumber("1 "))
	t.Log(isNumber("078332e437"))
}

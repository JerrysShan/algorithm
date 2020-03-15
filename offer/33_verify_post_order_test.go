package offer

import "testing"

func TestVerifyPostorder(t *testing.T) {
	t.Log(verifyPostorder([]int{1, 6, 3, 2, 5}))
	t.Log(verifyPostorder([]int{1, 3, 2, 6, 5}))
}

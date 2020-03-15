package offer

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	return verify(postorder, 0, len(postorder)-1)
}

func verify(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}
	index := start
	for ; index < end; index++ {
		if postorder[index] > postorder[end] {
			break
		}
	}
	for i := index; i < end; i++ {
		if postorder[i] < postorder[end] {
			return false
		}
	}
	return verify(postorder, start, index-1) && verify(postorder, index, end-1)
}

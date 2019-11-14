package leetcode

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(r1 *TreeNode, r2 *TreeNode) bool {
	if r1 == nil && r2 == nil {
		return true
	}
	if r1 == nil || r2 == nil {
		return false
	}

	return r1.Val == r2.Val && isMirror(r1.Left, r2.Right) && isMirror(r1.Right, r2.Left)
}

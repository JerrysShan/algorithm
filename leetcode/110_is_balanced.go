package leetcode

func isBalanced(root *TreeNode) bool {
	d := depth(root)
	return d >= 0
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left < 0 || right < 0 {
		return -1
	}
	if left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

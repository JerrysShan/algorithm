package leetcode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + depth(root.Left)
	right := 1 + depth(root.Right)
	if left > right {
		return left
	}
	return right
}

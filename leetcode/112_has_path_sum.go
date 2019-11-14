package leetcode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	temp := sum - root.Val
	if root.Right == nil && root.Left == nil {
		return temp == 0
	}

	return hasPathSum(root.Left, temp) || hasPathSum(root.Right, temp)
}

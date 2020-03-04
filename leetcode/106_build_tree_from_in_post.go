package leetcode

func buildTree2(inorder []int, postorder []int) *TreeNode {
	var postIndex = len(postorder) - 1
	idxMap := make(map[int]int)
	for k, v := range inorder {
		idxMap[v] = k
	}
	return helpBuildTree2(postorder, idxMap, &postIndex, 0, len(inorder))
}

func helpBuildTree2(postorder []int, idxMap map[int]int, postIndex *int, left, right int) *TreeNode {
	if left == right {
		return nil
	}
	root := &TreeNode{Val: postorder[*postIndex]}
	*postIndex--
	index := idxMap[root.Val]
	root.Right = helpBuildTree2(postorder, idxMap, postIndex, index+1, right)
	root.Left = helpBuildTree2(postorder, idxMap, postIndex, left, index)
	return root
}

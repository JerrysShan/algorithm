package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	var preIndex = 0
	idxMap := make(map[int]int)
	for k, v := range inorder {
		idxMap[v] = k
	}
	return helperBuildTree(preorder, idxMap, &preIndex, 0, len(inorder))
}

func helperBuildTree(preorder []int, idxMap map[int]int, preIndex *int, left, right int) *TreeNode {
	if left == right {
		return nil
	}
	root := &TreeNode{Val: preorder[*preIndex]}
	*preIndex++
	index := idxMap[root.Val]
	root.Left = helperBuildTree(preorder, idxMap, preIndex, left, index)
	root.Right = helperBuildTree(preorder, idxMap, preIndex, index+1, right)
	return root
}

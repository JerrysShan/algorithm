package leetcode

func constructFromPrePost(pre []int, post []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	if len(pre) == 1 {
		return &TreeNode{Val: pre[0]}
	}
	leftLength := 0
	for i := 0; i < len(pre); i++ {
		if post[i] == pre[1] {
			leftLength = i + 1
		}
	}
	root := &TreeNode{Val: pre[0]}
	root.Left = constructFromPrePost(pre[1:leftLength+1], post[0:leftLength])
	root.Right = constructFromPrePost(pre[leftLength+1:len(pre)], post[leftLength:len(pre)-1])
	return root
}

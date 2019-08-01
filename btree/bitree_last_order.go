package btree

import (
	"algorithm/common"
	"algorithm/structure"
	"fmt"
)

// BinTreeLastOrderRecursion 二叉树的后序遍历递归算法
func BinTreeLastOrderRecursion(node *common.Node) {
	if node == nil {
		return
	}
	BinTreeLastOrderRecursion(node.Left)
	BinTreeLastOrderRecursion(node.Right)
	fmt.Println(node.Data)
}

// BinTreeLastOrder 二叉树的后序遍历非递归算法
func BinTreeLastOrder(root *common.Node) {
	if root == nil {
		return
	}
	cur := root
	stack := &structure.Stack{}
	var pre *common.Node
	for cur != nil || stack.Size() > 0 {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		temp := stack.Top().(*common.Node)
		if temp.Right == nil || temp.Right == pre {
			fmt.Print(temp.Data, "->")
			pre = temp
			stack.Pop()
		} else {
			cur = temp.Right
		}
	}
}

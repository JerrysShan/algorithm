package btree

import (
	"algorithm/common"
	"fmt"
)

//二叉树的中序遍历递归算法
func BinTreeMiddleOrderRecursion(root *common.Node) {
	if root == nil {
		return
	}
	BinTreeMiddleOrderRecursion(root.Left)
	fmt.Print(root.Data, "->")
	BinTreeMiddleOrderRecursion(root.Right)
}

//二叉树的中序遍历非递归算法

func bitreeMiddleOrder() {

}

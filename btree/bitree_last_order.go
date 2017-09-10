package btree

import (
	"algorithm/common"
	"fmt"
)

//二叉树的后序遍历递归算法
func BinTreeLastOrderRecursion(node *common.Node) {
	if node == nil {
		return
	}
	BinTreeLastOrderRecursion(node.Left)
	BinTreeLastOrderRecursion(node.Right)
	fmt.Println(node.Data)
}

//二叉树的后序遍历非递归算法

func bitreeLastOrder() {

}

package btree

import "testing"

func TestBinTreeMiddleOrder(t *testing.T) {
	node1.Left = node2
	node1.Right = node5
	node2.Left = node3
	node2.Right = node4
	node5.Left = node6
	node5.Right = node7
	BinTreeMiddleOrder(node1)
}

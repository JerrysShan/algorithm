package search

import (
	"algorithm/btree"
	"algorithm/common"
	"fmt"
)

//BinaryTreeSearch 二叉查找树
type BinaryTreeSearch struct {
	Root *common.Node
}

//Show 按有序的方式展示这颗二叉查找树
func (bt *BinaryTreeSearch) Show() {
	btree.BinTreeMiddleOrderRecursion(bt.Root)
	fmt.Println()
}

//Size 获取树的大小
func (bt *BinaryTreeSearch) Size() int {
	return bt.size(bt.Root)
}

func (bt *BinaryTreeSearch) size(node *common.Node) int {
	if node == nil {
		return 0
	}
	return node.N
}

//Put 插入一个元素
func (bt *BinaryTreeSearch) Put(key int) {
	bt.Root = bt.put(key, bt.Root)
}

func (bt *BinaryTreeSearch) put(key int, node *common.Node) *common.Node {
	if node == nil {
		return &common.Node{
			Data: key,
			N:    1,
		}
	}
	if key < node.Data {
		node.Left = bt.put(key, node.Left)
	} else if key > node.Data {
		node.Right = bt.put(key, node.Right)
	} else {
		node.Data = key
	}
	node.N = bt.size(node.Left) + bt.size(node.Right) + 1
	return node
}

//Get 获取一个元素
func (bt *BinaryTreeSearch) Get(key int) *common.Node {
	return bt.get(key, bt.Root)
}

func (bt *BinaryTreeSearch) get(key int, node *common.Node) *common.Node {
	if node == nil {
		return nil
	}

	if key < node.Data {
		return bt.get(key, node.Left)
	} else if key > node.Data {
		return bt.get(key, node.Right)
	} else {
		return node
	}
}

//Min 获取最小值
func (bt *BinaryTreeSearch) Min() int {
	node := bt.min(bt.Root)
	return node.Data
}

func (bt *BinaryTreeSearch) min(node *common.Node) *common.Node {
	if node.Left == nil {
		return node
	}
	return bt.min(node.Left)
}

//Max 获取最大值
func (bt *BinaryTreeSearch) Max() int {
	node := bt.max(bt.Root)
	return node.Data
}
func (bt *BinaryTreeSearch) max(node *common.Node) *common.Node {
	if node.Right == nil {
		return node
	}
	return bt.max(node.Right)
}

//Floor 返回此数据结构中小于等于key的元素值
func (bt *BinaryTreeSearch) Floor(key int) *common.Node {
	return bt.floor(key, bt.Root)
}

func (bt *BinaryTreeSearch) floor(key int, node *common.Node) *common.Node {
	if node == nil {
		return nil
	}

	if key == node.Data {
		return node
	}
	if key < node.Data {
		return bt.floor(key, node.Left)
	}
	t := bt.floor(key, node.Right)
	if t != nil {
		return t
	}
	return node
}

//Ceiling 返回此数据结构中大于等于key的元素值
func (bt *BinaryTreeSearch) Ceiling(key int) *common.Node {
	return bt.ceiling(key, bt.Root)
}

func (bt *BinaryTreeSearch) ceiling(key int, node *common.Node) *common.Node {
	if node == nil {
		return nil
	}

	if key == node.Data {
		return node
	}
	if key > node.Data {
		return bt.ceiling(key, node.Right)
	}
	t := bt.ceiling(key, node.Left)
	if t != nil {
		return t
	}
	return node
}

// Rank 返回key的排名
func (bt *BinaryTreeSearch) Rank(key int) int {
	return bt.rank(key, bt.Root)
}

func (bt *BinaryTreeSearch) rank(key int, node *common.Node) int {
	if node == nil {
		return 0
	}
	if key < node.Data {
		return bt.rank(key, node.Left)
	} else if key > node.Data {
		return 1 + bt.size(node.Left) + bt.rank(key, node.Right)
	} else {
		return bt.size(node.Left)
	}
}

//Index 返回排名为key的节点
func (bt *BinaryTreeSearch) Index(rank int) *common.Node {
	return bt.index(rank, bt.Root)
}

func (bt *BinaryTreeSearch) index(rank int, node *common.Node) *common.Node {
	if node == nil {
		return nil
	}
	t := bt.size(node.Left)
	if t > rank {
		return bt.index(rank, node.Left)
	} else if t < rank {
		return bt.index(rank-t-1, node.Right)
	}
	return node
}

//DeleteMin 删除此数据结构中最小元素
func (bt *BinaryTreeSearch) DeleteMin() {
	bt.Root = bt.deleteMin(bt.Root)
}

func (bt *BinaryTreeSearch) deleteMin(node *common.Node) *common.Node {
	if node.Left == nil {
		return node.Right
	}
	node.Left = bt.deleteMin(node.Left)
	node.N = bt.size(node.Left) + bt.size(node.Right) + 1
	return node
}

//DeleteMax 删除此数据结构中最大元素
func (bt *BinaryTreeSearch) DeleteMax() *common.Node {
	return nil
}

//Delete 从此数据结构中删除一个元素
func (bt *BinaryTreeSearch) Delete(key int) {

}

func (bt *BinaryTreeSearch) delete(key int, node *common.Node) *common.Node {
	if node == nil {
		return nil
	}
	if key < node.Data {
		node.Left = bt.delete(key, node.Left)
	} else if key > node.Data {
		node.Right = bt.delete(key, node.Right)
	} else {
		if node.Right == nil {
			return node.Left
		}
		if node.Left == nil {
			return node.Right
		}
		t := node
		node = bt.min(t.Right)
		node.Right = bt.deleteMin(t.Right)
		node.Left = t.Left
	}
	node.N = bt.size(node.Left) + bt.size(node.Right) + 1
	return node
}

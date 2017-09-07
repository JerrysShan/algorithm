package search

import (
	"errors"
)

// BinarySearch only store integer that should grater than 0
type BinarySearch struct {
	Keys []int
	N    int
}

func (bs *BinarySearch) Size() int {
	return bs.N
}

// 非递归实现
// Rank方法具有以下性质:
// 1.如果该表中存在该键，rank返回该键的位置，也就是小于它的键的数量
// 2.如果表中不存在该键，rank还是应该返回表中的小于它的键的数量
func (bs *BinarySearch) Rank(key int) int {
	low := 0
	high := bs.N - 1
	for low <= high {
		mid := low + (high-low)/2
		if bs.Keys[mid] > key {
			high = mid - 1
		} else if bs.Keys[mid] < key {
			low = mid + 1
		} else {
			return mid
		}
	}
	return low
}

//递归实现
// func (bs *BinarySearch) Rank(key, low, high int) int {
// 	if high < low {
// 		return low
// 	}
// 	mid := low + (high-low)/2
// 	if bs.keys[mid] > key {
// 		return bs.Rank(key, low, mid-1)
// 	} else if bs.keys[mid] < key {
// 		return bs.Rank(key, mid+1, high)
// 	} else {
// 		return mid
// 	}
// }

//Get return key's index otherwise return -1
func (bs *BinarySearch) Get(key int) int {
	if bs.N == 0 {
		return -1
	}
	index := bs.Rank(key)
	if index < bs.N && bs.Keys[index] == key {
		return index
	}
	return -1
}

//Put will put key in the bs.Keys then bs.Keys should keep sorted ,if key <0 return an error
func (bs *BinarySearch) Put(key int) error {
	if key <= 0 {
		return errors.New("key should greater than 0")
	}
	i := bs.Rank(key)
	if i < bs.N && bs.Keys[i] == key {
		return nil
	}
	for j := bs.N; j > i; j-- {
		bs.Keys[j] = bs.Keys[j-1]
	}
	bs.Keys[i] = key
	bs.N++
	return nil
}

// Floor 函数 如果key 在这个列表中，则返回key，如果不在则返回小于这个key的最大元素,
//小于列表最小元素的key 不存在Floor
func (bs *BinarySearch) Floor(key int) int {
	i := bs.Rank(key)
	if i < bs.N && bs.Keys[i] == key {
		return bs.Keys[i]
	}
	if i == 0 {
		return 0
	}
	return bs.Keys[i-1]
}

// Ceiling 函数返回如果key在这个列表中，则返回key，如果不在则返回大于这个key的最小元素
//大于列表最大元素的key 不存在Ceiling
func (bs *BinarySearch) Ceiling(key int) int {
	index := bs.Rank(key)
	if index == bs.N {
		return 0
	}
	return bs.Keys[index]
}

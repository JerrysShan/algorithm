package search

var Total int

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
		Total++
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

func (bs *BinarySearch) get(key int) int {
	if bs.N == 0 {
		return -1
	}
	index := bs.Rank(key)
	if index < bs.N && bs.Keys[index] == key {
		return index
	}
	return -1
}

func (bs *BinarySearch) Put(key int) {
	i := bs.Rank(key)
	if i < bs.N && bs.Keys[i] == key {
		return
	}
	for j := bs.N; j > i; j-- {
		Total++
		bs.Keys[j] = bs.Keys[j-1]
	}
	bs.Keys[i] = key
	bs.N++
}

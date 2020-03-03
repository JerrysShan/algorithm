package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	arr := []int{}
	cur := head
	for cur != nil {
		arr = append(arr, head.Val)
		cur = cur.Next
	}
	i := 0
	j := len(arr) - 1
	for j > i {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

package leetcode

import "fmt"

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var arr []*ListNode
	p := head
	for p != nil {
		arr = append(arr, p)
		p = p.Next
	}
	length := len(arr)
	half := length / 2
	// if length%2 != 0 {
	// 	half = half + 1
	// }
	fmt.Println(half, length)
	for i := 0; i < half; i++ {
		temp := arr[i].Next
		arr[i].Next = arr[length-i-1]
		arr[length-i-1].Next = temp
	}
	arr[half].Next = nil
}

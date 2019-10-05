package leetcode

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p := head
	q := head.Next
	for p != nil && q != nil {
		if p == q {
			return true
		}
		p = p.Next
		if q.Next != nil {
			q = q.Next.Next
		} else {
			return false
		}
	}
	return false
}

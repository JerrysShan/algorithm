package leetcode

func middleNode(head *ListNode) *ListNode {
	one := head
	two := head
	for two != nil && two.Next != nil {
		one = one.Next
		two = two.Next.Next
	}
	return one
}

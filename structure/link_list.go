package structure

type Node struct {
	data interface{}
	next *Node
}

func IsLoop(head *Node) bool {
	if head == nil {
		return false
	}
	m := head
	n := head.next.next
	for {

		if m == nil || n == nil {
			break
		}
		if m == n {
			return true
		}
		m = m.next
		n = n.next.next
	}
	return false
}

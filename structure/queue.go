package structure

type Queue struct {
	head *nextNode
	tail *nextNode
	size int
}
type nextNode struct {
	data interface{}
	next *nextNode
}

func (q *Queue) EnQueue(d interface{}) {
	if q.head == nil {
		q.head = &nextNode{
			data: d,
		}
		q.tail = q.head
	} else {
		n := &nextNode{
			data: d,
		}
		q.tail.next = n
		q.tail = n
	}
	q.size++
}

func (q *Queue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	n := q.head
	q.head = n.next
	return n.data
}

func (q *Queue) Size() int {
	return q.size
}

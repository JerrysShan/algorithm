package structure

type Stack struct {
	top  *node
	head *node
	size int
}
type node struct {
	data interface{}
	pre  *node
}

func (s *Stack) Push(d interface{}) {
	if s.head == nil {
		s.head = &node{
			data: d,
		}
		s.top = s.head
	} else {
		n := &node{
			data: d,
		}
		n.pre = s.top
		s.top = n
	}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.top == nil {
		return nil
	}
	n := s.top
	s.top = n.pre
	return n.data
}

func (s *Stack) Size() int {
	return s.size
}

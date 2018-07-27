package structure

type SkipList struct {
	level int
	head  *SkipListNode
}

type SkipListNode struct {
	key   int
	value int
	next  []SkipListNode
}

func CreateNode(level, key, value int) *SkipListNode {
	return nil
}

func CreateList() *SkipList {
	return nil
}

func (s *SkipList) Insert(n *SkipListNode) {

}

func (s *SkipList) Delete(key int) {

}

func (s *SkipList) Search(key int) {

}

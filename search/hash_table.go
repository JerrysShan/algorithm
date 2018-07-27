package search

const SIZE = 16

type KV struct {
	key string
	val string
}

type LinkNode struct {
	Data KV
	Next *LinkNode
}

func NewLinkNode() *LinkNode {
	return &LinkNode{KV{"", ""}, nil}
}

func (node *LinkNode) AddNode(key, val string) {
	n := &LinkNode{KV{key, val}, nil}
	node.Next = n
}

func hashCode(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % SIZE
}

type HashTable struct {
	Size    int
	buckets [SIZE]*LinkNode
}

func NewHashTable() *HashTable {
	m := &HashTable{}
	for i := 0; i < SIZE; i++ {
		m.buckets[i] = NewLinkNode()
	}
	return m
}

func (h *HashTable) AddKV(key, val string) {
	index := hashCode(key)
	link := h.buckets[index]
	if link.Data.key == "" && link.Next == nil {
		link.Data.key = key
		link.Data.val = val
	} else {
		link.AddNode(key, val)
	}
}

func (h *HashTable) GetVal(key string) string {
	index := hashCode(key)
	link := h.buckets[index]
	head := link
	var value string
	for head != nil {
		if head.Data.key == key {
			value = head.Data.val
			break
		} else {
			head = head.Next
		}
	}
	return value
}

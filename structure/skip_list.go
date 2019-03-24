package structure

import "math/rand"

type SkipList struct {
	level int
	head  *SkipListNode
}

type SkipListNode struct {
	key   int
	value int
	next  []*SkipListNode
}

func random(n int) int {
	level := rand.Intn(n)
	for level == 0 {
		level = rand.Intn(n)
	}
	return level
}

func createNode(level, key, value int) *SkipListNode {
	return &SkipListNode{key: key, value: value, next: make([]*SkipListNode, random(level))}
}

func CreateList(maxLevel int) *SkipList {
	return &SkipList{level: maxLevel, head: &SkipListNode{next: make([]*SkipListNode, maxLevel)}}
}

func (s *SkipList) Insert(key, value int) bool {
	n := createNode(s.level, key, value)
	newNodeMaxLevel := len(n.next)
	cur := s.head
	for l := s.level - 1; l >= 0; l-- {
		if cur.next[l] == nil {
			if l < newNodeMaxLevel {
				n.next[l] = cur.next[l]
				cur.next[l] = n
			}
		}
		for cur.next[l] != nil && cur.next[l].key < key {
			cur = cur.next[l]
		}
		if cur.next[l] != nil && cur.next[l].key == key {
			return false
		}
		if l < newNodeMaxLevel {
			n.next[l] = cur.next[l]
			cur.next[l] = n
		}
	}
	return true
}

func (s *SkipList) Delete(key int) bool {
	cur := s.head
	var delete *SkipListNode
	for l := s.level; l >= 0; l-- {
		if cur.next[l] == nil {
			continue
		}
		for cur.next[l] != nil && cur.next[l].key < key {
			cur = cur.next[l]
		}
		if cur.next[l] != nil && key == cur.next[l].key {
			delete = cur.next[l]
			cur.next[l] = delete.next[l]
			delete.next[l] = nil
		}
	}
	if delete == nil {
		return false
	}
	return true
}

func (s *SkipList) Search(key int) interface{} {
	cur := s.head
	for l := s.level - 1; l >= 0; l-- {
		if cur.next[l] == nil {
			continue
		}
		for cur.next[l] != nil && cur.next[l].key < key {
			cur = cur.next[l]
		}
		if cur.next[l] != nil && cur.next[l].key == key {
			return cur.next[l].value
		}
	}
	return nil
}

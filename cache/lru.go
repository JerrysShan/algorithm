package cache

type lruNode struct {
	key  string
	val  interface{}
	next *lruNode
	pre  *lruNode
}

// LRU 近期最少使用
type LRU struct {
	head     *lruNode
	tail     *lruNode
	capacity int
	store    map[string]*lruNode
}

// NewLRU return LRU
func NewLRU(capacity int) *LRU {
	return &LRU{capacity: capacity, store: make(map[string]*lruNode)}
}

// Get return val by key
func (l *LRU) Get(key string) interface{} {
	node, ok := l.store[key]
	if ok {
		l.removeAndInsert(node)
		return node.val
	}
	return nil
}

// Put set key ,val
func (l *LRU) Put(key string, val interface{}) {
	if l.head == nil {
		l.head = &lruNode{key: key, val: val}
		l.tail = l.head
		l.store[key] = l.head
	}
	node := l.store[key]
	if node != nil {
		node.val = val
		l.removeAndInsert(node)
	} else {
		tmp := &lruNode{key: key, val: val}
		if len(l.store) > l.capacity {
			delete(l.store, l.tail.key)
		}
		l.store[key] = tmp
		tmp.next = l.head
		l.head.pre = tmp
		l.head = tmp
	}
}

func (l *LRU) removeAndInsert(node *lruNode) {
	if node == l.head {
		return
	}
	if node == l.tail {
		l.tail = l.tail.pre
		l.tail.next = nil
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
	// 插入到头结点
	node.next = l.head
	node.pre = nil
	l.head.pre = node
	l.head = node
}

package loadbalance

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// Hash 一致性算法,用于大规模缓存系统的负载均衡
type Hash struct {
	keys        []int
	nodes       map[int]string
	replication int
}

func hashCode(str string) int {
	hashCode := crc32.ChecksumIEEE([]byte(str))
	if hashCode < 0 {
		hashCode = -hashCode
	}
	return int(hashCode)
}

// NewHash return a init Hash
func NewHash(numberOfReplicas int) *Hash {
	hosts := make(map[int]string)
	indexs := []int{}
	return &Hash{
		keys:        indexs,
		nodes:       hosts,
		replication: numberOfReplicas,
	}
}

//Add 添加一个节点
func (h *Hash) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < h.replication; i++ {
			code := hashCode(node + strconv.Itoa(i))
			h.nodes[code] = node
			h.keys = append(h.keys, code)
		}
	}
	sort.Ints(h.keys)
}

// Get return node
func (h *Hash) Get(key string) string {
	if h.IsEmpty() {
		return ""
	}
	code := hashCode(key)
	index := sort.Search(len(h.keys), func(n int) bool {
		return h.keys[n] >= code
	})
	if index == len(h.keys) {
		index = 0
	}
	return h.nodes[h.keys[index]]
}

// IsEmpty return hash is empty
func (h *Hash) IsEmpty() bool {
	return len(h.keys) == 0
}

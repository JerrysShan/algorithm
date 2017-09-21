package search

const RED bool = true
const BLACK bool = false

type Node struct {
	data  int
	color bool
	left  *Node
	right *Node
}

type RedBlackBST struct {
	root *Node
}

func (rb *RedBlackBST) isRed(h *Node) bool {
	return h.color == RED
}

func (rb *RedBlackBST) rotateLeft(h *Node) *Node {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	return x
}

func (rb *RedBlackBST) rotateRight(h *Node) *Node {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	return x
}

func (rb *RedBlackBST) filpColors(h *Node) {
	h.left.color = BLACK
	h.right.color = BLACK
	h.color = RED
}

func (rb *RedBlackBST) putRoot(key int) {
	rb.root = rb.put(rb.root, key)
	rb.root.color = BLACK
}

func (rb *RedBlackBST) put(h *Node, key int) *Node {
	if h == nil {
		return &Node{
			data:  key,
			color: RED,
		}
	}
	if key < h.data {
		h.left = rb.put(h.left, key)
	} else if key > h.data {
		h.right = rb.put(h.right, key)
	}
	if rb.isRed(h.right) && !rb.isRed(h.left) {
		h = rb.rotateLeft(h)
	}

	if rb.isRed(h.left) && rb.isRed(h.left.left) {
		h = rb.rotateRight(h)
	}

	if rb.isRed(h.left) && rb.isRed(h.right) {
		rb.filpColors(h)
	}
	return h
}

func (rb *RedBlackBST) DeleteMin() {
	if !rb.isRed(rb.root) && !rb.isRed(rb.root) {
		rb.root.color = RED
	}
	rb.root = rb.deleteMin(rb.root)
	rb.root.color = BLACK
}

func (rb *RedBlackBST) deleteMin(node *Node) *Node {
	if node.left == nil {
		return nil
	}
	return nil
}

func (rb *RedBlackBST) moveRedLeft(node *Node) *Node {
	rb.filpColors(node)
	if rb.isRed(node.right.left) {
		node.right = rb.rotateRight(node.right)
		node = rb.rotateLeft(node)
		rb.filpColors(node)
	}
	return node
}

func (rb *RedBlackBST) moveRedRight(node *Node) *Node {
	rb.filpColors(node)
	if rb.isRed(node.left.left) {
		node = rb.rotateRight(node)
		rb.filpColors(node)
	}
	return node
}

func (rb *RedBlackBST) balance(h *Node) *Node {
	if rb.isRed(h.right) {
		h = rb.rotateLeft(h)
	}
	if rb.isRed(h.left) && rb.isRed(h.left.left) {
		h = rb.rotateRight(h)
	}
	if rb.isRed(h.left) && rb.isRed(h.right) {
		rb.filpColors(h)
	}

	// h.size = size(h.left) + size(h.right) + 1
	return h
}

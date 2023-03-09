package algods

import "golang.org/x/exp/constraints"

type binaryTreeNode[T constraints.Ordered] struct {
	data                T
	left, right, parent *binaryTreeNode[T]
}

type BinaryTree[T constraints.Ordered] struct {
	root  *binaryTreeNode[T]
	Count int
}

func (t *BinaryTree[T]) Add(value T) {
	t.Count++

	if t.root == nil {
		t.root = new(binaryTreeNode[T])
		t.root.data = value
		return
	}

	current := t.root

	for true {
		if value <= current.data {
			if current.left == nil {
				current.left = new(binaryTreeNode[T])
				current.left.data = value
				current.left.parent = current
				return
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.right = new(binaryTreeNode[T])
				current.right.data = value
				current.right.parent = current
				return
			} else {
				current = current.right
			}
		}
	}
}

func findNodeWithValue[T constraints.Ordered](t *binaryTreeNode[T], value T) (bool, *binaryTreeNode[T]) {
	if t == nil {
		return false, nil
	}

	if t.data == value {
		return true, t
	}

	if value < t.data {
		return findNodeWithValue(t.left, value)
	}

	return findNodeWithValue(t.right, value)
}

func (n *binaryTreeNode[T]) removeImmediateLeafChild(child *binaryTreeNode[T]) {
	if !child.isLeaf() {
		panic("Attempt to remove a non-leaf node from the tree")
	}

	if n.left == child {
		n.left = nil
		return
	}

	if n.right == child {
		n.right = nil
		return
	}

	panic("Attempt to remove a child node that isn't actually a child node")
}

func (n *binaryTreeNode[T]) rightMostLeftChild() *binaryTreeNode[T] {
	current := n.left
	for !current.isLeaf() {
		current = current.right
	}

	return current
}

func (n *binaryTreeNode[T]) leftMostRightChild() *binaryTreeNode[T] {
	current := n.right
	for !current.isLeaf() {
		current = current.left
	}

	return current
}

func (n *binaryTreeNode[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *binaryTreeNode[T]) isRoot() bool {
	return n.parent == nil
}

func (t *BinaryTree[T]) Remove(value T) bool {
	found, node := findNodeWithValue(t.root, value)

	if !found {
		return false
	}

	t.Count--

	if node.isLeaf() {
		if node.isRoot() {
			t.root = nil
			return true
		}

		node.parent.removeImmediateLeafChild(node)
		return true
	}

	if node.left != nil {
		replaceLeaf := node.rightMostLeftChild()
		node.data = replaceLeaf.data
		replaceLeaf.parent.removeImmediateLeafChild(replaceLeaf)
		return true
	}

	replaceLeaf := node.leftMostRightChild()
	node.data = replaceLeaf.data
	replaceLeaf.parent.removeImmediateLeafChild(replaceLeaf)
	return true
}

func (t *BinaryTree[T]) Contains(value T) bool {
	found, _ := findNodeWithValue(t.root, value)

	return found
}

func (t *BinaryTree[T]) PreOrder(callback func(value T)) {
	preOrderInternal(t.root, callback)
}

func preOrderInternal[T constraints.Ordered](node *binaryTreeNode[T], callback func(value T)) {
	if node == nil {
		return
	}

	callback(node.data)
	preOrderInternal(node.left, callback)
	preOrderInternal(node.right, callback)
}

func (t *BinaryTree[T]) InOrder(callback func(value T)) {
	inOrderInternal(t.root, callback)
}

func inOrderInternal[T constraints.Ordered](node *binaryTreeNode[T], callback func(value T)) {
	if node == nil {
		return
	}

	inOrderInternal(node.left, callback)
	callback(node.data)
	inOrderInternal(node.right, callback)
}

func (t *BinaryTree[T]) PostOrder(callback func(value T)) {
	postOrderInternal(t.root, callback)
}

func postOrderInternal[T constraints.Ordered](node *binaryTreeNode[T], callback func(value T)) {
	if node == nil {
		return
	}

	postOrderInternal(node.left, callback)
	postOrderInternal(node.right, callback)
	callback(node.data)
}

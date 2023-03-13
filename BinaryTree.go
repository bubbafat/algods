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
	newNode := new(binaryTreeNode[T])
	newNode.data = value

	for true {
		if value <= current.data {
			if current.left == nil {
				current.setLeftChild(newNode)
				return
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.setRightChild(newNode)
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

func (n *binaryTreeNode[T]) replaceImmediateChild(child *binaryTreeNode[T], replacement *binaryTreeNode[T]) {

	if n.left == child {
		n.setLeftChild(replacement)
		return
	}

	if n.right == child {
		n.setRightChild(replacement)
		return
	}

	panic("Attempt to replace a child node that isn't actually a child node")
}

func (n *binaryTreeNode[T]) rightMostLeftChild() *binaryTreeNode[T] {
	current := n.left
	for current.right != nil {
		current = current.right
	}

	return current
}

func (n *binaryTreeNode[T]) leftMostRightChild() *binaryTreeNode[T] {
	current := n.right
	for current.left != nil {
		current = current.left
	}

	return current
}

func (n *binaryTreeNode[T]) children() int {
	switch {
	case n.left == nil && n.right == nil:
		return 0
	case n.left != nil && n.right != nil:
		return 2
	default:
		return 1
	}
}

func (n *binaryTreeNode[T]) isLeaf() bool {
	if n == nil {
		panic("whoa")
	}

	return n.children() == 0
}

func (n *binaryTreeNode[T]) isRoot() bool {
	return n.parent == nil
}

func (n *binaryTreeNode[T]) setLeftChild(child *binaryTreeNode[T]) {
	n.left = child
	if child != nil {
		child.parent = n
	}
}

func (n *binaryTreeNode[T]) setRightChild(child *binaryTreeNode[T]) {
	n.right = child
	if child != nil {
		child.parent = n
	}
}

func (t *BinaryTree[T]) setRoot(child *binaryTreeNode[T]) {
	t.root = child
	if child != nil {
		child.parent = nil
	}
}

func (t *BinaryTree[T]) Remove(value T) bool {
	found, node := findNodeWithValue(t.root, value)

	if !found {
		return false
	}

	t.Count--

	// case 1: node has no right child, node.left replaces node
	if node.right == nil {
		if node.isRoot() {
			t.setRoot(node.left)
		} else {
			node.parent.replaceImmediateChild(node, node.left)
		}
		return true
	}

	// case 2: node's right child has not left child, node's right child replaces current
	if node.right.left == nil {
		node.right.setLeftChild(node.left)

		if node.isRoot() {
			t.setRoot(node.right)
		} else {
			node.parent.replaceImmediateChild(node, node.right)
		}
		return true
	}

	// case 3: node's right child has a left child. replace node node's right child's left-most child
	leftMost := node.leftMostRightChild()

	leftMost.parent.setLeftChild(leftMost.right)
	leftMost.setLeftChild(node.left)
	leftMost.setRightChild(node.right)

	if node.isRoot() {
		t.setRoot(leftMost)
	} else {
		node.parent.replaceImmediateChild(node, leftMost)
	}

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

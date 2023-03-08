package algods

import "golang.org/x/exp/constraints"

type binaryTreeNode[T constraints.Ordered] struct {
	data        T
	left, right *binaryTreeNode[T]
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
				return
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.right = new(binaryTreeNode[T])
				current.right.data = value
				return
			} else {
				current = current.right
			}
		}
	}
}

func (t *BinaryTree[T]) Remove(value T) bool {
	panic("not implemented")
}

func (t *BinaryTree[T]) Contains(value T) bool {
	panic("not implemented")
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

package tree

type AVLNode[T any] struct {
	Value T
	Left *AVLNode[T]
	Right *AVLNode[T]
	height int
}

type AVLTree[T any] struct {
	root *AVLNode[T]
}

func height[T any](node *AVLNode[T]) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (avl *AVLNode[T]) updateHeight() {
	avl.height = 1 + max(height(avl.Right), height(avl.Left))
}

func balanceFactor[T any](node *AVLNode[T]) int {
	if node == nil {
		return 0
	}
	return height(node.Right) - height(node.Left)
}
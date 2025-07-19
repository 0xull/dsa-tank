package tree

import "cmp"

type BST[T cmp.Ordered] struct {
	Root *LinkedBTNode[T]
}

func (bst *BST[T]) Search(val T) bool {
	return bst.searchRecursive(bst.Root, val)
}

func (bst *BST[T]) searchRecursive(node *LinkedBTNode[T], val T) bool {
	if node == nil {
		return false
	}
	
	if node.Value == val {
		return true
	} else if node.Value > val {
		return bst.searchRecursive(node.Left, val)
	} else {
		return bst.searchRecursive(node.Right, val)
	}
}

func (bst *BST[T]) SearchIterative(val T) bool {
	c := bst.Root
	
	for c != nil {
		if c.Value == val {
			return true
		}
		
		if c.Value > val {
			c = c.Left
		} else {
			c = c.Right
		}
	}
	
	return false
}

func (bst *BST[T]) Insert(val T) {
	bst.Root = bst.insertRecursive(bst.Root, val)
}

func (bst *BST[T]) insertRecursive(node *LinkedBTNode[T], val T) *LinkedBTNode[T] {
	if node == nil {
		return &LinkedBTNode[T]{Value: val }
	}
	
	if val < node.Value {
		node.Left = bst.insertRecursive(node.Left, val)
	} else if val > node.Value {
		node.Right = bst.insertRecursive(node.Right, val)
	}
	
	return node
}

func (bst *BST[T]) InsertIterative(node *LinkedBTNode[T], val T) {
	nn := &LinkedBTNode[T]{Value: val}
	
	if bst.Root == nil {
		bst.Root = nn
	}
	
	curr := node
	for {
		if val < curr.Value {
			if curr.Left == nil {
				curr.Left = nn
				return
			}
			curr = curr.Left
		} else if val > curr.Value {
			if curr.Right == nil {
				curr.Right = nn
				return
			}
			curr = curr.Right
		} else {
			return
		}
	}
}

func (bst *BST[T]) Delete(val T) {
	bst.Root = bst.delRecursive(bst.Root, val)
}

func (bst *BST[T]) delRecursive(node *LinkedBTNode[T], val T) *LinkedBTNode[T] {
	if node == nil {
		return nil
	}
	
	if val < node.Value {
		node.Left = bst.delRecursive(node.Left, val)
	} else if val > node.Value {
		node.Right = bst.delRecursive(node.Right, val)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		
		successor := findMin(node.Right)
		node.Value = successor.Value
		
		node.Right = bst.delRecursive(node.Right, successor.Value)
	}
	
	return node
}

func findMin[T any](node *LinkedBTNode[T]) *LinkedBTNode[T] {
	curr := node
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

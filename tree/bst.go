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
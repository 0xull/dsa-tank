package tree

import (
	"cmp"
)

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type RBNode[T cmp.Ordered] struct {
	Value T
	Color Color
	Left *RBNode[T]
	Right *RBNode[T]
	Parent *RBNode[T]
}

type RedBlackTree[T cmp.Ordered] struct {
	root *RBNode[T]
}

func (rbt *RedBlackTree[T]) Insert(val T) {
	nn := &RBNode[T]{Value: val, Color: RED}
	
	if rbt.root == nil {
		rbt.root = nn
	} else {
		curr := rbt.root
		var parent *RBNode[T]
		for curr != nil {
			parent = curr
			if val < curr.Value {
				curr = curr.Left
			} else if val > curr.Value {
				curr = curr.Right
			} else {
				return
			}
		}
		
		nn.Parent = parent
		if val < parent.Value {
			parent.Left = nn
		} else {
			parent.Right = nn
		}
	}
	
	rbt.fixInsert(nn)
}

func (rbt *RedBlackTree[T]) fixInsert(z *RBNode[T]) {
	for z != rbt.root && z.Parent.Color == RED {
		parent := z.Parent
		grandParent := parent.Parent
		
		if parent == grandParent.Left {
			uncle := grandParent.Right
			
			if uncle != nil && uncle.Color == RED {
				parent.Color = BLACK
				uncle.Color = BLACK
				grandParent.Color = RED
				z = grandParent
			} else {
				if z == parent.Right {
					z = parent
					rbt.leftRotate(z)
					parent = z.Parent
				}
				
				parent.Color = BLACK
				grandParent.Color = RED
				rbt.rightRotate(grandParent)
			}
		} else {
			uncle := grandParent.Left
			
			if uncle != nil && uncle.Color == RED {
				parent.Color = BLACK
				uncle.Color = BLACK
				grandParent.Color = RED
				z = grandParent
			} else {
				if z == parent.Left {
					z = parent
					rbt.rightRotate(z)
					parent = z.Parent
				}
				
				parent.Color = BLACK
				grandParent.Color = RED
				rbt.leftRotate(grandParent)
			}
		}
	}
	
	rbt.root.Color = BLACK
}

func (rbt *RedBlackTree[T]) leftRotate(node *RBNode[T]) {}
func (rbt *RedBlackTree[T]) rightRotate(node *RBNode[T]) {}

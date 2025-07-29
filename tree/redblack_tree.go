package tree

import "cmp"

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
	
}

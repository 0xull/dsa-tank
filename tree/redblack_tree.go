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
	Value  T
	Color  Color
	Left   *RBNode[T]
	Right  *RBNode[T]
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

func (rbt *RedBlackTree[T]) leftRotate(x *RBNode[T]) {
	y := x.Right
	x.Right = y.Left

	if y.Left != nil {
		x.Right.Parent = x
	}

	y.Parent = x.Parent
	if x.Parent == nil {
		rbt.root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func (rbt *RedBlackTree[T]) rightRotate(y *RBNode[T]) {
	x := y.Left
	y.Left = x.Right

	if x.Right != nil {
		y.Left.Parent = y
	}

	x.Parent = y.Parent
	if y.Parent == nil {
		rbt.root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	x.Right = y
	y.Parent = x
}

func (rbt *RedBlackTree[T]) Delete(val T) {
	nodeToDelete := rbt.findNode(val)
	if nodeToDelete == nil {
		return
	}
	rbt.deleteNode(nodeToDelete)
}

func (rbt *RedBlackTree[T]) deleteNode(z *RBNode[T]) {
	var y = z
	yOriginalColor := y.Color

	// x is the child that moves into y's original position
	var x *RBNode[T]

	// case 1 & 2: z has at most one child.
	if z.Left == nil {
		x = z.Right
		rbt.transplant(z, z.Right)
	} else if z.Right == nil {
		x = z.Left
		rbt.transplant(z, z.Left)
	} else {
		// case 3: z has two children
		// find in-order successor
		y = rbFindMin(z.Right)
		yOriginalColor = y.Color
		x = y.Right

		if y.Parent == z {
			if x != nil {
				x.Parent = y
			}
		} else {
			rbt.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		rbt.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		if x != nil {
			rbt.deleteFixup(x)
		}
	}
}

// deleteFixup restores the RedBlack properties after a deletion.
// 
// NOTE: sadly i don't fully have a grasp of this implementation yet. it is somewhat slippy.
// might have to setup an observable tree sample and explore the diff scenarios involved.
// that could possibly take awhile before done, but yeah, definitely.
func (rbt *RedBlackTree[T]) deleteFixup(x *RBNode[T]) {
	for x != rbt.root && x.Color == BLACK {
		if x == x.Parent.Left {
			w := x.Parent.Right
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.leftRotate(x.Parent)
				w = x.Parent.Right
			}
			if (w.Left == nil || w.Left.Color == BLACK) && (w.Right == nil || w.Right.Color == BLACK) {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right == nil || w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					rbt.rightRotate(w)
					w = x.Parent.Right
				}
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				rbt.leftRotate(x.Parent)
				x = rbt.root
			}
		} else {
			w := x.Parent.Left
			if w.Color == RED {
				w.Color = BLACK
				x.Parent.Color = RED
				rbt.rightRotate(x.Parent)
				w = x.Parent.Left
			}

			if (w.Left == nil || w.Left.Color == BLACK) && (w.Right == nil || w.Right.Color == BLACK) {
				w.Color = RED
				x = x.Parent
			} else {
				if w.Left == nil || w.Left.Color == BLACK {
					w.Right.Color = BLACK
					w.Color = RED
					rbt.leftRotate(w)
					w = x.Parent.Left
				}

				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				rbt.rightRotate(x.Parent)
				x = rbt.root
			}
		}
	}

	x.Color = BLACK
}

// transplant replaces the subtree rooted at node u with the subtree rooted at node v
// It handles all the necessary parent pointer updates.
func (rbt *RedBlackTree[T]) transplant(u, v *RBNode[T]) {
	if u.Parent == nil {
		rbt.root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

func (rbt *RedBlackTree[T]) findNode(val T) *RBNode[T] {
	node := rbt.root
	for node != nil && node.Value != val {
		if val < node.Value {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}

func rbFindMin[T cmp.Ordered](node *RBNode[T]) *RBNode[T] {
	curr := node
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

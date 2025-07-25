package tree

import "cmp"

type AVLNode[T cmp.Ordered] struct {
	Value  T
	Left   *AVLNode[T]
	Right  *AVLNode[T]
	height int
}

type AVLTree[T cmp.Ordered] struct {
	root *AVLNode[T]
}

func height[T cmp.Ordered](node *AVLNode[T]) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (avl *AVLNode[T]) updateHeight() {
	avl.height = 1 + max(height(avl.Right), height(avl.Left))
}

func balanceFactor[T cmp.Ordered](node *AVLNode[T]) int {
	if node == nil {
		return 0
	}
	return height(node.Right) - height(node.Left)
}

// rightRotate performs a right rotation on the pivot node y.
// It is called to fix a Left-Left imbalance.
//
// The function is triggered by this transition:
//
// Initial Balanced State  -->  Insert 10...  -->  Imbalanced State (Input to this function)
//
//	  30 (y)                                         30 (y)  <-- BF = -2
//	 /                                              /
//	20 (x)                                         20 (x)
//	                                              /
//	                                             10 (T1)
//
// The rotation transforms the imbalanced input into the final, balanced output.
//
// Imbalanced Input (y is 30)      Balanced Output (returns x, which is 20)
//
//	    30 (y)                            20 (x)
//	   /                                 /  \
//	  20 (x)         -->                10   30
//	 /
//	10 (T1)
func rightRotate[T cmp.Ordered](y *AVLNode[T]) *AVLNode[T] {
	// 1. Identify the key nodes from the imbalanced input.
	//    - y is the pivot, node 30.
	//    - x becomes its left child, node 20.
	x := y.Left
	//    - T2 is the right child of x. In our example, node 20 has
	//      no right child, so T2 is assigned the value nil.
	T2 := x.Right

	// 2. Perform the rotation by reassigning the pointers.
	//    The right child of x (20) becomes y (30).
	x.Right = y
	//    The left child of y (30) becomes T2 (which is nil in this case).
	y.Left = T2

	// 3. Update the heights of the nodes whose children have changed.
	//    We MUST update the former pivot (y) first, as it is now a child.
	y.updateHeight() // y's children are now both nil. height = 1 + max(-1,-1) = 0.
	x.updateHeight() // x's children are now 10 (h=0) and 30 (h=0). height = 1 + max(0,0) = 1.

	// 4. Return the pointer to the new root of this subtree (node 20).
	return x
}

// leftRotate performs a left rotation on the pivot node x.
// It is called to fix a Right-Right imbalance.
//
// The function is triggered by this transition:
//
// Initial Balanced State  -->  Insert 30...  -->  Imbalanced State (Input to this function)
//
//	10 (x)                                       10 (x)  <-- BF = +2
//	  \                                            \
//	   20 (y)                                       20 (y)
//	                                                  \
//	                                                   30 (T3)
//
// The rotation transforms the imbalanced input into the final, balanced output.
//
// Imbalanced Input (x is 10)      Balanced Output (returns y, which is 20)
//
//	10 (x)                              20 (y)
//	  \                                /  \
//	   20 (y)      -->                10   30
//	     \
//	      30 (T3)
func leftRotate[T cmp.Ordered](x *AVLNode[T]) *AVLNode[T] {
	// 1. Identify key nodes from the imbalanced input.
	//    - x is the pivot, node 10.
	//    - y becomes its right child, node 20.
	y := x.Right
	//    - T2 is the left child of y. In this example, node 20
	//      has no left child, so T2 is assigned nil.
	T2 := y.Left

	// 2. Perform the rotation.
	//    The left child of y (20) becomes x (10).
	y.Left = x
	//    The right child of x (10) becomes T2 (nil).
	x.Right = T2

	// 3. Update heights. We must update the former pivot (x) first.
	x.updateHeight() // x's children are now both nil. height = 0.
	y.updateHeight() // y's children are 10 (h=0) and 30 (h=0). height = 1.

	// 4. Return the pointer to the new root of this subtree (node 20).
	return y
}

func (avl *AVLTree[T]) Insert(value T) {
	avl.root = avl.insert(avl.root, value)
}

func (avl *AVLTree[T]) insert(node *AVLNode[T], val T) *AVLNode[T] {
	if node == nil {
		return &AVLNode[T]{Value: val, height: 0}
	}

	if val == node.Value {
		return node
	}

	if val < node.Value {
		node.Left = avl.insert(node.Left, val)
	} else if val > node.Value {
		node.Right = avl.insert(node.Right, val)
	}

	node.updateHeight()
	bf := balanceFactor(node)

	if bf > 1 {
		if val > node.Right.Value {
			return leftRotate(node)
		}
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}
	if bf < -1 {
		if val < node.Left.Value {
			return rightRotate(node)
		}
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	return node
}

func (avl *AVLTree[T]) Delete(value T) {
	avl.root = avl.delete(avl.root, value)
}

func (avl *AVLTree[T]) delete(node *AVLNode[T], val T) *AVLNode[T] {
	if node == nil {
		return nil
	}

	if val < node.Value {
		node.Left = avl.delete(node.Left, val)
	} else if val > node.Value {
		node.Right = avl.delete(node.Right, val)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		successor := avlFindMin(node.Right)
		node.Value = successor.Value
		node.Right = avl.delete(node.Right, successor.Value)
	}

	node.updateHeight()
	bf := balanceFactor(node)

	if bf > 1 {
		if balanceFactor(node.Right) >= 0 {
			return leftRotate(node)
		}
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}
	if bf < -1 {
		if balanceFactor(node.Left) <= 0 {
			return rightRotate(node)
		}
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	return node
}

func avlFindMin[T cmp.Ordered](node *AVLNode[T]) *AVLNode[T] {
	curr := node
	if curr != nil && curr.Left == nil {
		curr = curr.Left
	}
	return curr
}

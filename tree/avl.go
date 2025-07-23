package tree

type AVLNode[T comparable] struct {
	Value  T
	Left   *AVLNode[T]
	Right  *AVLNode[T]
	height int
}

type AVLTree[T comparable] struct {
	root *AVLNode[T]
}

func height[T comparable](node *AVLNode[T]) int {
	if node == nil {
		return -1
	}
	return node.height
}

func (avl *AVLNode[T]) updateHeight() {
	avl.height = 1 + max(height(avl.Right), height(avl.Left))
}

func balanceFactor[T comparable](node *AVLNode[T]) int {
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
func rightRotate[T comparable](y *AVLNode[T]) *AVLNode[T] {
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
func leftRotate[T comparable](x *AVLNode[T]) *AVLNode[T] {
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

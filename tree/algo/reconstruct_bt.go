package bt_algo

import "github.com/IkehAkinyemi/dsa-tank/tree"


func ReconstructBTTreeFromPreorderInorder[T comparable](preorder, inorder []T) *tree.LinkedBTNode[T] {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Index preorder for O(1) lookup
	pLookup := make(map[T]int)
	for i, v := range inorder {
		pLookup[v] = i
	}

	preorderIndex := 0

	var builderHelper func(leftIndex, rightIndex int  ) *tree.LinkedBTNode[T]
	builderHelper = func(leftIndex, rightIndex int,) *tree.LinkedBTNode[T] {
		if leftIndex > rightIndex {
			return nil
		}

		node := &tree.LinkedBTNode[T]{Value: preorder[preorderIndex]}
		preorderIndex++

		nodeIndex := pLookup[node.Value]

		node.Left = builderHelper(leftIndex , nodeIndex-1)
		node.Right = builderHelper(nodeIndex + 1, rightIndex)

		return node
	}

	return builderHelper(0, len(inorder)-1)
}

func ReconstructBTTreeFromPostorderInorder[T comparable](postorder, inorder []T) *tree.LinkedBTNode[T] {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	mLookup := make(map[T]int)
	for i, v := range inorder {
		mLookup[v] = i
	}

	pIndex := len(postorder)-1
	var builder func(lIndex, rIndex int) *tree.LinkedBTNode[T]
	builder = func(lIndex, rIndex int) *tree.LinkedBTNode[T] {
		if lIndex > rIndex {
			return nil
		}

		node := &tree.LinkedBTNode[T]{Value: postorder[pIndex]}
		pIndex--

		inIndex := mLookup[node.Value]

		node.Right = builder((inIndex+1), rIndex)
		node.Left = builder(lIndex, (inIndex-1))
		
		return node
	}

	return builder(0, len(inorder)-1)
}

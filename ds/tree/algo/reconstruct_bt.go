package bt_algo

import "github.com/IkehAkinyemi/dsa-tank/ds/tree"

func ReconstructBTTree[T comparable](preorder, inorder []T) *tree.LinkedBTNode[T] {
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
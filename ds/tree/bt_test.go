package tree_test

import (
	"fmt"
	"testing"

	"github.com/IkehAkinyemi/dsa-tank/ds/tree"
)

func TestBTNodeTraversal(t *testing.T) {
	// Make tree:
	// Let's build the following tree:
	//	      F
	//	     / \
	//	    B   G
	//	   / \   \
	//	  A   D   I
	//	     / \
	//	    C   E
	root := &tree.LinkedBTNode[string]{Value: "F"}
	root.Right = &tree.LinkedBTNode[string]{Value: "G"}
	root.Right.Right = &tree.LinkedBTNode[string]{Value: "I"}
	root.Left = &tree.LinkedBTNode[string]{Value: "B"}
	root.Left.Left = &tree.LinkedBTNode[string]{Value: "A"}
	root.Left.Right = &tree.LinkedBTNode[string]{Value: "D"}
	root.Left.Right.Left = &tree.LinkedBTNode[string]{Value: "C"}
	root.Left.Right.Right = &tree.LinkedBTNode[string]{Value: "E"}
	
	// Recursive traversal
	fmt.Print("Recursive traversal: \ngot: ")
	tree.PreOrderRecursive(root)
	fmt.Print("\nexpected: F B A D C E G I \n\n")
	
	// Iterative traversal
	fmt.Printf("Iterative traversal: \ngot: ")
	tree.PreOrderIterative(root)
	fmt.Print("\nexpected: F B A D C E G I \n")
}
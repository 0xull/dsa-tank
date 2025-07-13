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
	
	// Recursive Pre-order traversal
	fmt.Print("Recursive NLR traversal: \ngot: ")
	tree.PreOrderRecursive(root)
	fmt.Print("\nexpected: F B A D C E G I \n\n")
	
	// Iterative Pre-order traversal
	fmt.Printf("Iterative NLR traversal: \ngot: ")
	tree.PreOrderIterative(root)
	fmt.Print("\nexpected: F B A D C E G I \n\n")
	
	// Recursive In-order traversal
	fmt.Print("Recursive LNR traversal: \not: ")
	tree.InOrderRecursive(root)
	fmt.Print("\nexpected: A B C D E F G I \n\n")
	
	// Iterative In-order traversal
	fmt.Print("Iterative LNR traversal: \ngot: ")
	tree.InOrderIterative(root)
	fmt.Print("\nexpected: A B C D E F G I \n")
}
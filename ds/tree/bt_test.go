package tree_test

import (
	"fmt"
	"testing"

	"github.com/IkehAkinyemi/dsa-tank/ds/tree"
)

func sampleBTNode() *tree.LinkedBTNode[string] {
	// Make tree:
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
	
	return root
}

func TestNLRTraversal(t *testing.T) {
	root := sampleBTNode()
	
	// Recursive Pre-order traversal
	fmt.Print("Recursive NLR traversal: \ngot: ")
	tree.PreOrderRecursive(root)
	fmt.Print("\nexpected: F B A D C E G I \n\n")
	
	// Iterative Pre-order traversal
	fmt.Printf("Iterative NLR traversal: \ngot: ")
	tree.PreOrderIterative(root)
	fmt.Print("\nexpected: F B A D C E G I \n\n")
}

func TestLNRTraversal(t *testing.T) {
	root := sampleBTNode()
	
	// Recursive In-order traversal
	fmt.Print("Recursive LNR traversal: \not: ")
	tree.InOrderRecursive(root)
	fmt.Print("\nexpected: A B C D E F G I \n\n")
	
	// Iterative In-order traversal
	fmt.Print("Iterative LNR traversal: \ngot: ")
	tree.InOrderIterative(root)
	fmt.Print("\nexpected: A B C D E F G I \n\n")
}

func TestLRNTraversal(t *testing.T) {
	root := sampleBTNode()
	
	// Recursive Post-order traversal
	fmt.Print("Recursive LRN traversal: \ngot: ")
	tree.PostOrderRecursive(root)
	fmt.Print("\nexpected: A C E D B I G F \n\n")
	
	// Iterative Post-order traversal 
	fmt.Print("Iterative LRN traversal: \ngot: ")
	tree.PostOrderIterative(root)
	fmt.Print("\nexpected: A C E D B I G F \n")
}
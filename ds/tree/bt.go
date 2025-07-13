package tree

import "fmt"

// LinkedBTNode reps a linked binary tree in memory with
// n >= 0 nodes, but n <= 2 nodes.
type LinkedBTNode[T any] struct {
	Value T
	Left *LinkedBTNode[T]
	Right *LinkedBTNode[T]
}

type ArrayBT[T any] []T

// PreOrderRecursive performs a recursive pre-order traversal on the tree in 
// NLR systematic pattern.
func PreOrderRecursive[T any](node *LinkedBTNode[T]) {
	if node == nil {
		return
	}
	
	// process node
	fmt.Printf("%v ", node.Value)
	
	// Recursive: traverse the left child node
	PreOrderRecursive(node.Left)
	
	// Recursive: traverse the right child node
	PreOrderRecursive(node.Right)
}

// PreOrderIterative performs a pre-order traversal using iterative stack method
// in a NLR systematic pattern.
func PreOrderIterative[T any](node *LinkedBTNode[T]) {
	if node == nil {
		return
	}
	
	stack := []*LinkedBTNode[T]{node}
	
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// process node value
		fmt.Printf("%v ", node.Value)
		
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	
} 

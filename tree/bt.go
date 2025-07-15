package tree

import (
	"fmt"
)

// LinkedBTNode reps a linked binary tree in memory with
// n >= 0 nodes, but n <= 2 nodes.
type LinkedBTNode[T any] struct {
	Value T
	Left  *LinkedBTNode[T]
	Right *LinkedBTNode[T]
}

type ArrayBT[T any] []T

// PreOrderRecursive performs a recursive pre-order traversal on a tree in
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
// in NLR systematic pattern.
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

// InOrderIterative perform iterative in-order traversal of a tree in 
// LNR systematic pattern.
func InOrderIterative[T any](node *LinkedBTNode[T]) {
	if node == nil {
		return
	}

	stack := []*LinkedBTNode[T]{}
	current := node

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// process the node
		fmt.Printf("%v ", node.Value)
		
		current = node.Right
	}
}

// InOrderRecursive perform recursive in-order traversal of a tree in 
// LNR systematic pattern.
func InOrderRecursive[T any](node *LinkedBTNode[T]) {
	if node == nil {
		return
	}
	
	InOrderRecursive(node.Left)
	
	fmt.Printf("%v ", node.Value)
	
	InOrderRecursive(node.Right)
}

// PostOrderIterative performs an iterative post-order traversal on a
// tree using LRN systematic pattern. To achieve that, perform a skewed
// variant of NLR, that is NRL which is inverse of LRN.
func PostOrderIterative[T any](node *LinkedBTNode[T]) {
	if node == nil {
		return
	}
	
	stack := []*LinkedBTNode[T]{node}
	var result []T
	
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// process node
		result = append([]T{node.Value}, result...)
		
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	
	for _, v := range result {
		fmt.Printf("%v ", v)
	}
}

// PostOrderRecursive perform a recursive post-order traversal on a tree
// in a LRN systematic pattern.
func PostOrderRecursive[T any](node *LinkedBTNode[T]) {
	if node == nil {
		return
	}
	
	// Left subtree traversal
	PostOrderRecursive(node.Left)
	
	// Right subtree traversal
	PostOrderRecursive(node.Right)
	
	// Process node value
	fmt.Printf("%v ", node.Value)
}

// LevelOrderTraversal is a "breadth first search" traversal on a tree
func LevelOrderTraversal[T any](node *LinkedBTNode[T]) {
	if node == nil {
		return
	}
	
	queue := []*LinkedBTNode[T]{node}
	
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		
		fmt.Printf("%v ", node.Value)
		
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
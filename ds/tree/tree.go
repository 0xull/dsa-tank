package tree

// GeneralTree defined a tree data structure with unstricted amount
// of children nodes.
type GeneralTree[T any] struct {
	Value T
	Children []*GeneralTree[T]
}

// // LinkedBTNode reps a linked binary tree in memory with
// // n >= 0 nodes, but n <= 2 nodes.
// type LinkedBTNode[T any] struct {
// 	Value T
// 	Left *LinkedBTNode[T]
// 	Right *LinkedBTNode[T]
// }

// // ArrayBT is an array implementation of the BT data structure
// // with ops:
// //   parent node: (i - 1) / 2 (integer divisin, floor the obtained value)
// type ArrayBT[T any] []T
package tree

import "cmp"

type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

type RBNode[T cmp.Ordered] struct {
	Value T
	Color Color
	Left *RBNode[T]
	Right *RBNode[T]
	Parent *RBNode[T]
}

type RBTree[T cmp.Ordered] struct {
	root *RBNode[T]
}

package tree

import (
	"cmp"
	"errors"
)

type BTreeNode[T cmp.Ordered] struct {
	isLeaf   bool
	keys     []T
	children []*BTreeNode[T]

	// n is the current number of keys in the node.
	n int
}

type BTree[T cmp.Ordered] struct {
	root *BTreeNode[T]

	// t is the B-tree's minimum degree
	t int
}

// NewBTree initializes a new B-Tree.
func NewBTree[T cmp.Ordered](t int) (*BTree[T], error) {
	if t < 2 {
		return nil, errors.New("B-Tree must have a minimum degree of at least 2")
	}

	return &BTree[T]{
		root: &BTreeNode[T]{
			isLeaf:   true,
			keys:     make([]T, 2*t-1),
			children: make([]*BTreeNode[T], 2*t),
		},
		t: t,
	}, nil
}

func (btree *BTree[T]) Search(k T) (*BTreeNode[T], int) {
	if btree.root == nil {
		return nil, -1
	}
	
	btree.root.search(k)
}

func (x *BTreeNode[T]) search(k T) (*BTreeNode[T], int) {
	i := 0
	for i < x.n && k > x.keys[i] {
		i++
	}
	
	if i < x.n && k == x.keys[i] {
		return x, i
	}
	
	if x.isLeaf {
		return nil, -1
	}
	
	return x.children[i].search(k)
}

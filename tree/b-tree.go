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

	return btree.root.search(k)
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

func (btree *BTree[T]) Insert(k T) {
	root := btree.root
	if root.n == 2*btree.t-1 {
		newRoot := &BTreeNode[T]{isLeaf: false, children: make([]*BTreeNode[T], 2*btree.t), n: 0}
		newRoot.children[0] = root
		btree.root = newRoot
		btree.splitChild(newRoot, 0)
		btree.insertNonFull(newRoot, k)
	} else {
		btree.insertNonFull(root, k)
	}
}

// insertNonFull inserts a key into a node that is guaranteed to be not full.
func (btree *BTree[T]) insertNonFull(x *BTreeNode[T], k T) {
	i := x.n - 1
	if x.isLeaf {
		for i >= 0 && k < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = k
		x.n++
	} else {
		for i >= 0 && k < x.keys[i] {
			i--
		}
		i++
		if x.children[i].n == 2*btree.t-1 {
			btree.splitChild(x, i)
			if k > x.keys[i] {
				i++
			}
		}
		btree.insertNonFull(x.children[i], k)
	}
}

// splitChild splits the full child y of node x
func (btree *BTree[T]) splitChild(x *BTreeNode[T], i int) {
	y := x.children[i]
	z := &BTreeNode[T]{
		isLeaf:   y.isLeaf,
		keys:     make([]T, 2*btree.t-1),
		children: make([]*BTreeNode[T], 2*btree.t),
		n:        btree.t - 1,
	}

	for j := 0; j < btree.t-1; j++ {
		z.keys[j] = y.keys[j+btree.t]
	}

	if !y.isLeaf {
		for j := 0; j < btree.t; j++ {
			z.children[j] = y.children[j+btree.t]
		}
	}

	y.n = btree.t - 1
	for j := x.n; j > i; j-- {
		x.children[j+1] = x.children[j]
	}
	x.children[i+1] = z

	for j := x.n - 1; j >= i; j-- {
		x.keys[j+1] = x.keys[j]
	}
	x.keys[i] = y.keys[btree.t-1]

	x.n++
}
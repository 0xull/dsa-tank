package tree_test

import (
	"fmt"
	"testing"

	"github.com/IkehAkinyemi/dsa-tank/tree"
)

func TestBSTSearch(t *testing.T) {
	bst := &tree.BST[string]{
		Root: sampleBTNode(),
	}

	testCases := []struct {
		name  string
		input string
		exp   bool
	}{
		{
			name:  "a node within the tree",
			input: "F",
			exp:   true,
		},
	}

	for _, tc := range testCases {
		if got := bst.Search(tc.input); got != tc.exp {
			t.Errorf("expected: %v; got: %v", tc.exp, got)
		}

		if got := bst.SearchIterative(tc.input); got != tc.exp {
			t.Errorf("expected: %v; got: %v", tc.exp, got)
		}
	}
}

func TestBSTInsertion(t *testing.T) {
	vals := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	var bst tree.BST[int]

	for _, v := range vals {
		bst.Insert(v)
	}

	fmt.Print("got: ")
	tree.InOrderIterative(bst.Root)
	fmt.Printf("\nexpected: 1 3 4 6 7 8 10 13 14\n\n")

	bst.Insert(5)

	fmt.Print("got: ")
	tree.InOrderIterative(bst.Root)
	fmt.Printf("\nexpected: 1 3 4 5 6 7 8 10 13 14\n\n")
}

func TestBSTDeletion(t *testing.T) {
	vals := []int{6, 3, 8, 1, 0, 2, 4, 5, 7, 14, 12, 10, 9, 11, 13, 21, 19, 17, 15, 16, 18, 20, 24}
	var bst tree.BST[int]
	for _, v := range vals {
		bst.Insert(v)
	}
		
	bst.Delete(14)
	
	fmt.Print("got: ")
	tree.InOrderRecursive(bst.Root)
	fmt.Print("\nexpected: 0 1 2 3 4 5 6 7 8 9 10 11 12 13 15 16 17 18 19 20 21 24\n")
}

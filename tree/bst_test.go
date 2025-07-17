package tree_test

import (
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
			name: "a node within the tree",
			input: "F",
			exp: true,
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

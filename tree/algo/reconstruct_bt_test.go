package bt_algo

import (
	"reflect"
	"testing"

	"github.com/IkehAkinyemi/dsa-tank/tree"
)

func TestReconstructBTFromPreorderInorder(t *testing.T) {
	// Make tree:
	//	      F
	//	     / \
	//	    B   G
	//	   / \   \
	//	  A   D   I
	//	     / \
	//	    C   E
	expected := &tree.LinkedBTNode[string]{Value: "F"}
	expected.Right = &tree.LinkedBTNode[string]{Value: "G"}
	expected.Right.Right = &tree.LinkedBTNode[string]{Value: "I"}
	expected.Left = &tree.LinkedBTNode[string]{Value: "B"}
	expected.Left.Left = &tree.LinkedBTNode[string]{Value: "A"}
	expected.Left.Right = &tree.LinkedBTNode[string]{Value: "D"}
	expected.Left.Right.Left = &tree.LinkedBTNode[string]{Value: "C"}
	expected.Left.Right.Right = &tree.LinkedBTNode[string]{Value: "E"}

	preorder := []string{"F", "B", "A", "D", "C", "E", "G", "I"}
	inorder := []string{"A", "B", "C", "D", "E", "F", "G", "I"}
	
	got := ReconstructBTTreeFromPreorderInorder(preorder, inorder)
	
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %#v;\n\ngot: %#v;\n", expected, got)
	}

	postorder := []string{"A", "C", "E", "D", "B", "I", "G", "F"}
	got = ReconstructBTTreeFromPostorderInorder(postorder, inorder)
	
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %#v; \n\ngot: %#v;\n", expected, got)
	}
}

package bt_algo

import "testing"

func TestExpressionTree(t *testing.T) {
	// corresponds to (3 + 5) * (10 - 4)
	postfix := "3 5 + 10 4 - *"
	var expected float64 = 48 // 8 * 6
	
	root, err := ExpressionTree(postfix)
	if err != nil {
		t.Errorf("%#v", err)
	}
	
	got, err := EvaluateExpressionTree(root)
	if err != nil {
		t.Errorf("%#v", err)
	}
	
	if got != expected {
		t.Errorf("expected: %2.f; got: %2.f", expected, got)
	}
}
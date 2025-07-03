package recursion

import "testing"

func TestIterRecursionWithStack(t *testing.T) {
	exp := 120
	got := RecursionWithStack(5)
	if got != exp {
		t.Errorf("expected: %d; got: %d", exp, got)
	}
}
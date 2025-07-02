package recursion

import "testing"

func TestGCD(t *testing.T) {
	exp := 1
	if got := GCD(9, 16); got != exp {
		t.Errorf("expected: %d; got: %d", exp, got)
	}
}
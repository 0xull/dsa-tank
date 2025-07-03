package recursion

import "testing"

func TestIterFiboncci(t *testing.T) {
	exp := 120
	got := IterFactorial(5)
	if exp != got {
		t.Errorf("expected: %d; got: %d", exp, got)
	}
}
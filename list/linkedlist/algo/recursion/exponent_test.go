package recursion

import "testing"

func TestPower(t *testing.T) {
	exp := 81
	if got := Power(3, 4); got != exp {
		t.Errorf("expected: %d; got: %d", exp, got)
	}
}
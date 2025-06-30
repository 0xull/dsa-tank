package recursion_test

import (
	"testing"

	"github.com/IkehAkinyemi/dsa-tank/algos/recursion"
)

func TestFactorial(t *testing.T) {
	got := recursion.Factorial(4)
	var exp uint = 24
	if exp != got {
		t.Errorf("exp: %d; got: %d", exp, got)
	}
}
package recursion_test

import (
	"testing"

	"github.com/IkehAkinyemi/dsa-tank/algos/recursion"
)

func TestFactorial(t *testing.T) {
	got := recursion.Factorial(5)
	var exp uint = 120
	if exp != got {
		t.Errorf("exp: %d; got: %d", exp, got)
	}
}
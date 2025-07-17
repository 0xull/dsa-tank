package recursion_test

import (
	"testing"

	"github.com/IkehAkinyemi/dsa-tank/list/linkedlist/algo/recursion"
)

func TestMemoization(t *testing.T) {
	cache := make(map[int]int)
	exp := 21
	got := recursion.Memoization(8, cache)
	if got != exp {
		t.Errorf("expected: %d; got: %d", exp, got)
	}
}
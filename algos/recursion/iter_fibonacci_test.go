package recursion

import (
	"reflect"
	"testing"
)

func TestIterFibonacci(t *testing.T) {
	exp := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	if got := IterFibonacci(8); !reflect.DeepEqual(got, exp) {
		t.Errorf("expected: %d; got: %d", exp, got)
	}
}
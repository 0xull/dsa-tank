package leetcode

import (
	"math"
	"testing"
)

func reverse(x int) int {
	if x > -10 && x < 10 {
		return x
	}

	placeVal := 1
	isPos := true
	rev := 0

	if x < 0 {
		isPos = false
		x *= -1
	}

	for math.Pow10(placeVal) <= float64(x) {
		placeVal++
	}

	for i := placeVal-1; i >= 0; i-- {
		rev += (x % 10) * int(math.Pow10(i))
		x /= 10
	}

	if rev < math.MinInt32 || rev > math.MaxInt32 {
		return 0
	}

	if isPos {
		return rev
	}

	return rev * -1
}

func TestReverseInt(t *testing.T) {
	testCases := []struct {
		name   string
		args   int
		expect int
	}{{"[REVERSE INTEGER] first test", 123, 321},
		{"[REVERSE INTEGER] second test", -120, -21},
		{"[REVERSE INTEGER] third test", 20, 2},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := reverse(tc.args)
			if got != tc.expect {
				t.Errorf("got %d; expected %d", got, tc.expect)
			}
		})
	}
}

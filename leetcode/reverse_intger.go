package leetcode

import (
	"testing"
	"math"
)

func reverse(x int) int {
	if x > -10 && x < 10 {
        return x
    }

    placeVal := 0
    isPos := true
    rev := 0

    if x < 0 {
        isPos = false
        x *= -1
    }

    for math.Pow10(placeVal) < float64(x) {
        placeVal++
    }

    for i := placeVal; i >= 0; i-- {
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
	
}
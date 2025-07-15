package recursion

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	exp := 0
	if got := Fibonacci(3); got != exp {
		fmt.Println(got)
	}
}
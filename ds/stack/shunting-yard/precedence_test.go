package shuntingyard

import (
	"testing"
)

func TestComparePrecedence(t *testing.T) {
	testCases := []struct{
		op1 string
		op2 string
		exp int
		name string
	}{
		{"+", "+", 0, "add and add have same precedence"},
		{"+", "-", 0, "add and subtract have same precedence"},
		{"*", "%", 0, "multiply and modulus have same precedence"},
		{"*", "-", 1, "multiply has higher precedence than subtract"},
		{"/", "+", 1, "divide has higher precedence than add"},
		{"+", "*", -1, "add has lower precedence than multiply"},
		{"-", "%", -1, "subtract has lower precedence than modulus"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := comparePrecedence(tc.op1, tc.op2); got != tc.exp {
				t.Errorf("expected precedence comparison to be %d, but got %d", tc.exp, got)
			}
		})
	}
}
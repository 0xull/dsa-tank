package shuntingyard

import (
	"testing"
)

func TestEvalPrefix(t *testing.T) {
	testCases := []struct{
		name string
		input string
		exp string
		err error
	}{
		{
			name: "simple prefix expression",
			input: "+-92*/843",
			exp: "13",
			err: nil,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := EvalPrefix(tc.input); got != tc.exp {
				t.Errorf("expected: %s; got: %s", tc.exp, got)
			} else if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("failed with unexpected error: %v", err)
			}
		})
	}
}
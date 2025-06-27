package shuntingyard

import "testing"

func TestInfix2Prefix(t *testing.T) {
	testCases := []struct{
		name string
		input string
		exp string
		err error
	}{
		{
			name: "simple infix expression",
			input: "(9-2) + ((12/4)*3)",
			exp: "+-92*/1243",
			err: nil,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := Infix2Prefix(tc.input); got != tc.exp {
				t.Errorf("exp: %s; got: %s", tc.exp, got)
			} else if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("unexpected error: %s", err.Error())
			}
		})
	}
}
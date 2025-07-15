package shuntingyard

import (
	"fmt"
	"testing"
)

func TestShuntingYard(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		exp   string
		err   error
	}{
		{
			name:  "end-to-end spaced-parenthesized infix expression",
			input: "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )",
			exp:   "123+45**+",
			err:   nil,
		},
		{
			name:  "doubled middle parenthesized infix expression",
			input: "5+((1+2)*4)-3",
			exp:   "512+4*+3-",
			err:   nil,
		},
		{
			name:  "middle spaced-parenthesized infix expression",
			input: "3 * ( 4 + 2 ) - 6 / 3",
			exp:   "342+*63/-",
			err:   nil,
		},
		{
			name:  "simple infix expression",
			input: "3+4",
			exp:   fmt.Sprint("34+"),
			err:   nil,
		},
		{
			name:  "another simple infix expression",
			input: "1 + 4 * 6",
			exp:   "146*+",
			err:   nil,
		},
		{
			name:  "invalid infix expression (missing a opening parenthesis)",
			input: "1 + 3 - 4(5 -))",
			exp:   "",
			err:   fmt.Errorf("mismatched parenthese: no '(' found for its ')' at index 14"),
		},
		{
			name:  "invalid infix expression (single closing parenthesis)",
			input: ")",
			exp:   "",
			err:   fmt.Errorf("mismatched parenthese: no '(' found for its ')' at index 0"),
		},
		{
			name:  "invalid infix expression (single opening parenthesis)",
			input: "(",
			exp:   "",
			err:   fmt.Errorf("mismatched parenthese: no ')' found for '('"),
		},
		{
			name:  "invalid infix expression (missing a closing parenthesis)",
			input: "(1+4*6",
			exp:   "",
			err:   fmt.Errorf("mismatched parenthese: no ')' found for '('"),
		},
		{
			name:  "complex infix expression",
			input: "4 - (7 / 8 + (3 % 2 * 6) / 5) * 1",
			exp:   "478/32%6*5/+1*-",
			err:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if postfix, err := ShuntingYard(tc.input); postfix != tc.exp {
				t.Errorf("expected '%s'; got '%s'", tc.exp, postfix)
			} else if err != nil && tc.err.Error() != err.Error() {
				t.Errorf("%s", err.Error())
			}
		})
	}
}

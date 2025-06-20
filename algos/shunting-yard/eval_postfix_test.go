package shuntingyard

import "testing"

func TestPostfixEvaluation(t *testing.T) {
	testCases := []struct{
		name string
		input string
		exp string
		err error
	} {
		{
			name: "simple postfix expression",
			input: "123+45**+",
			exp: "101",
			err: nil,
		},
		{
			 name: "complex postfix expression",
			input: "482/73%6*5/+1*-", // manually check the arithmetics here
			exp: "",
			err: nil,
		},
	}
	t.Log(testCases[1])
}
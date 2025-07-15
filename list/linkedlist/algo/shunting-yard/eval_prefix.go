package shuntingyard

import (
	"fmt"
	"unicode"

	stack "github.com/IkehAkinyemi/dsa-tank/stack/example"
)

func EvalPrefix(prefix string) (string, error) {
	var ops stack.StackArray[string]
	r := []rune(prefix)
	
	for i := len(r)-1; i >= 0; i-- {
		switch {
		case unicode.IsLetter(r[i]) || unicode.IsDigit(r[i]):
			// if ops.Size() > 1 {
			// 	ops.Push(string(r[i]) + ops.Pop())
			// }
			ops.Push(string(r[i]))
		case isOperator(r[i]):
			op1 := ops.Pop()
			op2 := ops.Pop()
			
			r, err := evaluate(op2, op1, r[i])
			if err != nil {
				return "", err
			}
			ops.Push(r)
		}
	}
	
	if ops.Size() != 1 {
		return "", fmt.Errorf("invalid prefix expression: expected one resultant operator; got %d", ops.Size())
	}
	return ops.Pop(), nil
}
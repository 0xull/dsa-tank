package shuntingyard

import (
	"fmt"
	"strconv"

	stack "github.com/IkehAkinyemi/dsa-tank/ds/stack/example"
)

func EvalPostfix(postfix string) (string, error) {
	es := stack.NewStackArray[string](3)
	
	for _, c := range postfix {
		switch {
		case isOperand(c):
			es.Push(string(c))
		case isOperator(c):
			r, err := evaluate(es.Pop(), es.Pop(), c)
			if err != nil {
				return "", err
			}
			es.Push(r)
		}
	}
	
	if es.Size() != 1 {
		return "", fmt.Errorf("invalid postfix expression; stack size is %d after evaluation, expected 1", es.Size())
	}
	
	return es.Pop(), nil
}

func evaluate(operand1, operand2 string, operator rune) (string, error) {
	num1, err1 := strconv.ParseFloat(operand1, 64)
	num2, err2 := strconv.ParseFloat(operand2, 64)
	if err1 != nil || err2 != nil {
		return "", fmt.Errorf("failed ascii to integer conversion")
	}
	
	switch operator {
	case '+':
		return fmt.Sprint(num2 + num1), nil
	case '-':
		return fmt.Sprint(num2 - num1), nil
	case '*':
		return fmt.Sprint(num2 * num1), nil
	case '%':
		return fmt.Sprint(float64(int(num2) % int(num1))), nil
	case '/':
		if num1 == 0 {
			return "", fmt.Errorf("division by zero")
		}
		return fmt.Sprint(num2 / num1), nil
	default:
		return "", fmt.Errorf("operator not recognized")
	}
}
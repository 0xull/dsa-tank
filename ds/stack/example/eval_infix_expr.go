/* Check out https://mathcenter.oxford.emory.edu/site/cs171/evaluatingFullyParenthesizedInfixExpressions/ */
package stack

import (
	"fmt"
	"regexp"
	"strconv"
)

func EvalParenthesizedInfixExpression(expr string) string {
	sOperand := NewStackArray[string](len(expr)/2)
	sOperator := NewStackArray[string](len(expr)/2)
	
	for _, v := range expr {
		switch {
		case v == '(': continue
		case isOperator(v):
			sOperator.Push(string(v))
		case isOperand(v):
			sOperand.Push(string(v))
		case v == ')':
			num1 := sOperand.Pop()
			num2 := sOperand.Pop()
			operator := sOperator.Pop()
			
			result := evaluate(num1, num2, operator)
			sOperand.Push(result)
		}
	}
	
	return sOperand.Pop()
}

func isOperator(char rune) bool {
	switch char {
	case '+', '-', '*', '/':
		return true
	}
	return false
}

func isOperand(char rune) bool {
	matched, _ := regexp.Match("^[0-9]$", []byte(string(char)))
	return matched
}

func evaluate(num1, num2, operator string) string {
	op1, _ := strconv.Atoi(string(num1)) 
	op2, _ := strconv.Atoi(string(num2))
	switch operator {
	case "+":
		return fmt.Sprint((op2 + op1))
	case "-":
		return fmt.Sprint((op2 - op1))
	case "/":
		return fmt.Sprint((op2 / op1))
	case "*":
		return fmt.Sprint((op2 * op1))
	}
	return "0"
}

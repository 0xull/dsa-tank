package shuntingyard

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	stack "github.com/IkehAkinyemi/dsa-tank/ds/stack/example"
)

func ShuntingYard(expression string) (string, error) {
	s := stack.NewStackArray[rune](len(expression))
	var postfix strings.Builder

	s.Push('(')
	expression += ")"

	for i, v := range expression {
		switch {
		case isOperator(v):
			for {
				op := s.Peek()
				r := comparePrecedence(string(v), string(op))
				if op == '(' || r == 1 {
					break
				}
				postfix.WriteRune(s.Pop())
			}
			s.Push(v)
		case unicode.IsLetter(v) || unicode.IsDigit(v):
			postfix.WriteRune(v)
		case v == '(':
			s.Push(v)
		case v == ')':
			for !s.IsEmpty() && s.Peek() != '(' {
				postfix.WriteRune(s.Pop())
			}
			if s.IsEmpty() || s.Pop() != '(' {
				return "", fmt.Errorf("mismatched parenthese: no '(' found for its ')' at index %d", i-1)
			}
		}
	}

	if !s.IsEmpty() {
		return "", fmt.Errorf("mismatched parenthese: no ')' found for '('")
	}

	return postfix.String(), nil
}

func isOperator(char rune) bool {
	switch char {
	case '+', '-', '*', '/', '%':
		return true
	default:
		return false
	}
}

func isOperand(char rune) bool {
	matched, _ := regexp.Match("^[0-9]$", []byte(string(char)))
	return matched
}

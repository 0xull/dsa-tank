package shuntingyard

import (
	"fmt"
	"regexp"
	"strings"

	stack "github.com/IkehAkinyemi/dsa-tank/ds/stack/example"
)

func ShuntingYard(expression string) (string, error) {
	s := stack.NewStackArray[rune](len(expression))
	var postfix strings.Builder
	
	for i, v := range expression {
		switch {
		case isOperator(v):
			op := s.Peek()
			r := comparePrecedence(string(v), string(op))
			if r == 0 {
				// check for associativity
				if strings.Contains(lvl2Ops, string(op)) { // Left associativity
					postfix.WriteRune(op)
					s.Pop()
					s.Push(v)
				} else if strings.Contains(lvl3Ops, string(op)) { // Right associativity
					postfix.WriteRune(v)
				}
			} else if r == 1 {
				postfix.WriteRune(v)
			} else if r == -1 {
				s.Push(v)
			}
		case isOperand(v):
			postfix.WriteRune(v)
		case v == '(':
			s.Push(v)
		case v == ')':
			for !s.IsEmpty() && s.Peek() != '(' {
				postfix.WriteRune(s.Pop())
			}
			if !s.IsEmpty() && s.Pop() != '(' {
				return "", fmt.Errorf("mismatched parenthese: no '(' found for its ')' at index %d", i)
			}
		}
	}
	return postfix.String(), nil
}

func isOperator(char rune) bool {
	switch char {
	case '+', '-', '*', '/':
		return true
	default: return false
	}
}

func isOperand(char rune) bool {
	matched, _ := regexp.Match("^[0-9]$", []byte(string(char)))
	return matched
}

func associativity(substr string) string {
	if strings.Contains(lvl2Ops, substr) {
		return "left"
	} else if strings.Contains(lvl3Ops,  substr) {
		return "right"
	}
}
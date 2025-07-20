package bt_algo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ExprNode defines a node in the expression tree
type ExprNode struct {
	Value      string
	IsOperator bool
	LeftNode   *ExprNode
	RightNode      *ExprNode
}

func ExpressionTree(postfix string) (*ExprNode, error) {
	stack := []*ExprNode{}
	tokens := strings.Fields(postfix)

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				return nil, fmt.Errorf("invalid postfix expression, less than two operands for operator %s", token)
			}
			// Pop left & right operands
			rNode := stack[len(stack)-1]
			lNode := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			stack = append(stack, &ExprNode{Value: token, IsOperator: true, LeftNode: lNode, RightNode: rNode})
		} else {
			stack = append(stack, &ExprNode{Value: token, IsOperator: false})
		}
	}

	if len(stack) != 1 {
		return nil, errors.New("invalid postfix expression; final stack size is not 1")
	}

	return stack[0], nil
}

func isOperator(operator string) bool {
	return operator == "+" || operator == "-" || operator == "*" || operator == "/"
}

func EvaluateExpressionTree(node *ExprNode) (float64, error) {
	if node == nil {
		return 0, errors.New("empty expression tree")
	}

	if !node.IsOperator {
		val, err := strconv.ParseFloat(node.Value, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse node operand %s", node.Value)
		}
		return val, nil
	}

	lOperand, err := EvaluateExpressionTree(node.LeftNode)
	if err != nil {
		return 0, err
	}
	rOperand, err := EvaluateExpressionTree(node.RightNode)
	if err != nil {
		return 0, err
	}

	switch node.Value {
	case "+":
		return (lOperand + rOperand), nil
	case "-":
		return (lOperand - rOperand), nil
	case "*":
		return (lOperand * rOperand), nil
	case "/":
		if rOperand == 0 {
			return 0, errors.New("zero division not permitted")
		}
		return (lOperand / rOperand), nil
	default:
		return 0, fmt.Errorf("invalid operator value %s", node.Value)
	}
}

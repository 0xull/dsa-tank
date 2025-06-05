package shuntingyard

import "strings"

const (
	lvl3Ops = "*/%"
	lvl2Ops = "+-"
)

func comparePrecedence(op1 string, op2 string) int {
	if op1 == op2 || same(lvl3Ops, op1, op2) || same(lvl2Ops, op1, op2) {
		return 0
	}
	
	if strings.Contains(lvl3Ops, op1) {
		return 1
	}
	
	return -1
}

func same(lvl string, op1 string, op2 string) bool {
	return strings.Contains(lvl, string(op1)) && strings.Contains(lvl, op2)
}

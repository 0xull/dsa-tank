package shuntingyard

import "strings"

/*
 * Infix2Prefix is summarized in the below steps:
 * 	1. Reverse the given expression
 * 	2. Swap the ocurrence of ')' for '(' and vice versa.
 * 	3. Call a postfix algorithm (or writing one if none)
 * `4. Reverse the resulting string and return
 */

func Infix2Prefix(expression string) (string, error) {
	revExpr := reverseStr(expression)
	
	replr := strings.NewReplacer(")", "(", "(", ")")
	revExpr = replr.Replace(revExpr)
	
	pfx, err := ShuntingYard(revExpr)
	if err != nil {
		return "", err
	}
	
	return reverseStr(pfx), nil
}

func reverseStr(original string) string {
	r := []rune(original)
	for i, j := 0, len(original)-1; i < j; i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	
	return string(r)
}
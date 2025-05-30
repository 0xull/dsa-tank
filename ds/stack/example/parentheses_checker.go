package stack

func ValidateExpression(expr string) bool {
	stack := NewLinkedStack[rune](len(expr))
	
	for _, v := range expr {
		switch v {
		case '{', '[', '(':
			stack.Push(v) 
		case '}', ']', ')':
			if !stack.IsEmpty() {
				if !isMatchingPair(stack.Pop().elem, v) {
					return false
				}
				continue
			} else {
				return false
			}
		default: continue
		}
	}
	
	return stack.IsEmpty()
}

func isMatchingPair(open, close rune) bool {
	if close == ')' {
		return open == '('
	} 
	if close  == ']' {
		return open == '['
	}
	if close == '}' {
		return open == '{'
	}
	
	return false
}
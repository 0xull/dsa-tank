package recursion

import "github.com/IkehAkinyemi/dsa-tank/stack/example"

func RecursionWithStack(num int) int {
	if num <= 0 {
		return 0
	}
	
	var stack stack.StackArray[int]
	for n := 2; n <= num; n++ {
		stack.Push(n)
	}
	
	result := 1
	for !stack.IsEmpty() {
		result *= stack.Pop()
	}
	
	return result
}
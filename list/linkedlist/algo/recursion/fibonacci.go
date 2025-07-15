package recursion

func Fibonacci(number int) int {
	if number <= 1 {
		return number
	}
	
	return Fibonacci(number-1) + Fibonacci(number-2) 
}
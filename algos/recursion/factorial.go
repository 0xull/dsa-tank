package recursion

func Factorial(number uint) uint {
	if number <= 1 {
		return 1
	}
	return (number * Factorial(number - 1))
}
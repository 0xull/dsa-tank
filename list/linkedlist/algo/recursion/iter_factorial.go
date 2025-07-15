package recursion

func IterFactorial(number int) int {
	result := 1
	
	for n := 2; n <= number; n++ {
		result *= n
	}
	
	return result
}


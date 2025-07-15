package recursion

// Power calculates the base raised to the non-negative exponent
func Power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	
	return base * Power(base, exponent - 1)
}
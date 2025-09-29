package searchingandsorting

func CountingSort(slice []int) []int {
	n := len(slice)
	if n < 2 {
		return slice
	}

	max := 0
	for _, val := range slice {
		if val > max {
			max = val
		}
	}

	counts := make([]int, max+1)
	for _, val := range slice {
		counts[val]++
	}

	for i := 1; i <= max; i++ {
		counts[i] += counts[i-1]
	}

	output := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		val := slice[i]
		position := counts[val] - 1
		output[position] = val
		counts[val]--
	}
	
	return output
}

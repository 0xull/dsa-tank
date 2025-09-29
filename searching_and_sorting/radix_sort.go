package searchingandsorting

func RadixSort(slice []int) {
	n := len(slice)
	if n < 2 {
		return
	}

	max := 0
	for _, val := range slice {
		if val > max {
			max = val
		}
	}

	for exp := 1; max/exp > 0; exp *= 10 {
		countingSortForRadix(slice, exp)
	}
}

func countingSortForRadix(slice []int, exp int) {
	n := len(slice)
	output := make([]int, n)

	// the count array for digits (0-9)
	counts := make([]int, 10)

	for i := range n {
		index := (slice[i] / exp) % 10
		counts[index]++
	}

	for i := 1; i < 10; i++ {
		counts[i] += counts[i-1]
	}
	
	for i := n - 1; i >= 0; i-- {
		val := slice[i]
		index := (slice[i] / exp) % 10
		position := counts[index] - 1
		output[position] = val
		counts[index]--
	}
	
	copy(slice, output)
}

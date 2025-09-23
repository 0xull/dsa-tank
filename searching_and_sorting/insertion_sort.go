package searchingandsorting

import "cmp"

func InsertionSort[T cmp.Ordered](slice []T) {
	n := len(slice)
	if n < 2 {
		return
	}

	for i := 1; i < n; i++ {
		key := slice[i]

		j := i - 1
		
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		
		slice[j+1] = key
	}
}

package searchingandsorting

import "cmp"

func SelectionSort[T cmp.Ordered](slice []T) {
	n := len(slice)
	if n < 2 {
		return
	}
	
	for i := range n-1 {
		midIndex := i 
		
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[midIndex] {
				midIndex = j
			}
		}
		
		slice[i], slice[midIndex] = slice[midIndex], slice[i]
	}
}
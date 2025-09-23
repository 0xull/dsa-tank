package searchingandsorting

import "cmp"

func BubbleSort[T cmp.Ordered](slice []T) {
	n := len(slice)
	if n < 2 {
		return
	}
	
	for i := range n {
		swapped := false
		
		for j := range n-1-i {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
} 
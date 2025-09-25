package searchingandsorting

import "cmp"

func QuickSort[T cmp.Ordered](slice []T) {
	quickSortRec(slice, 0, len(slice)-1)
}

func quickSortRec[T cmp.Ordered](slice []T, low, high int) {
	if low < high {
		pivotIndex := partition(slice, low, high)

		quickSortRec(slice, low, pivotIndex-1)
		quickSortRec(slice, pivotIndex+1, high)
	}
}

func partition[T cmp.Ordered](slice []T, low, high int) int {
	pivot := slice[high]

	// 'i' is the index for the end of the "less than pivot" section
	i := low - 1

	for j := low; j < high; j++ {
		if slice[j] <= pivot {
			i++
			slice[i], slice[j] = slice[j], slice[i]
		}
	}

	slice[i+1], slice[high] = slice[high], slice[i+1]
	return i + 1
}

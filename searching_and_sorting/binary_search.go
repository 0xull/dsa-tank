package searchingandsorting

import "cmp"

func BinarySearch[T cmp.Ordered](slice []T, target T) (int, bool) {
	low := 0
	high := len(slice) - 1

	for low <= high {
		mid := low + (high-low)/2

		if slice[mid] == target {
			return mid, true
		} else if slice[mid] > target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	
	return -1, false
}

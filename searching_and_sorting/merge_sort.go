package searchingandsorting

import "cmp"

func MergeSort[T cmp.Ordered](slice []T) []T {
	n := len(slice)
	
	if n < 2 {
		return slice
	}
	
	mid := n / 2
	leftHalf := MergeSort(slice[:mid])
	rightHalf := MergeSort(slice[mid:])
	
	return merge(leftHalf, rightHalf)
}

func merge[T cmp.Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))
	
	i, j, k := 0, 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}
	
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}
	
	return result
}
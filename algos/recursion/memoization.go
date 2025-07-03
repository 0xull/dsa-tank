package recursion

func Memoization(num int, cache map[int]int) int {
	if v, ok := cache[num]; ok {
		return v
	}
	
	if num <= 1 {
		return num
	}
	
	result := Memoization(num-1, cache) + Memoization(num-2, cache)
	
	cache[num] = result
	return result
} 
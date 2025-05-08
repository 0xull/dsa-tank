package main

import "fmt"

// hanoi has the recurrence relation, T(n) = 2T(n-1) + 1, and its running time
// is T(n) = O(2^n)
func hanoi(n int, src, dst, tmp *[]string) {
	// Base case
	if n == 1 {
		*dst = append(*dst, (*src)[len(*src)-1])
		*src = (*src)[:len(*src)-1]
		return
	}
	
	// move first disk to destination
	hanoi(n-1, src, tmp, dst)
	
	// move second disk to destination
	*dst = append(*dst, (*src)[len(*src)-1])
	*src = (*src)[:len(*src)-1]
	
	// move temporary back to source
	hanoi(n-1, tmp, dst, src)
}

func main() {
	src := []string{"3", "2", "1"}
	dst := []string{}
	tmp := []string{}
	
	hanoi(len(src), &src, &dst, &tmp)
	
	fmt.Printf("Source peg: %+v", src)
	fmt.Printf("Destination peg: %+v", dst)
	fmt.Printf("Temporary peg: %+v", tmp)
}
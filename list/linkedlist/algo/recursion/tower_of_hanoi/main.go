package main

import "fmt"

func hanoi(n int, src, dst, tmp *[]string) {
	// Base case: If only one disk, move it directly from src to dst
	if n == 1 {
		*dst = append(*dst, (*src)[len(*src)-1])
		*src = (*src)[:len(*src)-1]
		fmt.Printf("Move disk from src to dst\n")
		return
	}
	// Recursive case: Move the top n-1 disks to tmp using dst as auxiliary
	hanoi(n-1, src, tmp, dst)

	// Move the nth (largest) disk from src to dst
	*dst = append(*dst, (*src)[len(*src)-1])
	*src = (*src)[:len(*src)-1]
	fmt.Printf("Move disk from src to dst\n")

	// Move the n-1 disks from tmp to dst using src as auxiliary
	hanoi(n-1, tmp, dst, src)
}

func main() {
	// Initialize the source peg with disks numbered from 3 to 1
	src := []string{"3", "2", "1", "0"}
	// Temporary and destination pegs start empty
	tmp := []string{}
	dst := []string{}

	// Call the Hanoi function to solve the puzzle
	hanoi(len(src), &src, &dst, &tmp)

	// Print the final state of the pegs
	fmt.Printf("Final State - Source: %+v\n", src)
	fmt.Printf("Final State - Destination: %+v\n", dst)
	fmt.Printf("Final State - Temporary: %+v\n", tmp)
}

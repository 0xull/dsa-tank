package recursion

func IterFibonacciSeq(number int) []int {
	seq := []int{0, 1}
	
	for i := 2; i <= number; i++ {
		seq = append(seq, seq[i-1] + seq[i-2])
	}
	
	return seq
}
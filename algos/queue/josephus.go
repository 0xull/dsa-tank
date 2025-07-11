package queue

import "fmt"

func JosephusProblem(n, k int) (int, error) {
	if n < 1 || k < 1 {
		return 0, fmt.Errorf("number of player or elimination sequence can't be zero")
	}
	if n == 1 {
		return 1, nil
	}

	queue := make([]int, n)
	for i := range n {
		queue[i] = i + 1
	}

	for len(queue) > 1 {
		for range k - 1 {
			queue = append(queue[1:], queue[0])
		}
		queue = queue[1:]
	}
	
	return queue[0], nil
}

package main

import (
	"fmt"
)

func main() {
	var n int
	var tmp int64

	fmt.Scanf("%d", &n)

	dims := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		dims[i] = tmp
	}

	m := 'A'
	for i := 1; i < len(dims); i++ {
		fmt.Printf("* %c: %dx%d\n", int(m)+i-1, dims[i-1], dims[i])
	}

	s, count := matrixChainOrder(dims)

	fmt.Printf("Optimal multiplication order: %s\nThe cost is: %d\n", s, count)
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
)

// Matrix Ai has dimensions dims[i-1] x dims[i]
// for i = 1 ... n
func matrixChainOrder(dims []int64) (string, int64) {
	if dims == nil || len(dims) < 3 {
		return "", 0
	}

	n := len(dims)

	m := make([][]int64, n)
	s := make([][]int64, n)

	for i := 0; i < n; i++ {
		m[i] = make([]int64, n)
		s[i] = make([]int64, n)
	}

	for l := 2; l < n; l++ {
		for i := 1; i < n-l+1; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt64
			for k := i; k <= j-1; k++ {
				cost := m[i][k] + m[k+1][j] + dims[i-1]*dims[k]*dims[j]
				if cost < m[i][j] {
					m[i][j] = cost
					// Each entry bracket[i,j]=k shows
					// where to split the product array
					// i,i+1....j for the minimum cost.
					s[i][j] = int64(k)
				}
			}
		}
	}

	// Prints the result in a bytes.Buffer
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)

	matrix := 'A'

	printResult(writer, 1, int64(len(dims)-1), s, &matrix)

	writer.Flush()
	return b.String(), m[1][n-1]
}

func printResult(buf io.Writer, i, j int64, s [][]int64, name *rune) {
	if i == j {
		fmt.Fprintf(buf, "%c", *name)
		*name = rune(int(*name) + 1)
		return
	}

	fmt.Fprintf(buf, "(")

	printResult(buf, i, s[i][j], s, name)
	printResult(buf, s[i][j]+1, j, s, name)

	fmt.Fprint(buf, ")")
}

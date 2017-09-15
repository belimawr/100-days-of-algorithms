package main

func multiplyMatrix(A, B [][]int) [][]int {
	n := len(A)
	m := len(A[0]) // Let's assume that the matrices are the correct size

	m = len(B)
	p := len(B[0])

	// Initialize the return matrix
	C := make([][]int, n)

	for i := 0; i < n; i++ {
		C[i] = make([]int, p)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			sum := 0
			for k := 0; k < m; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}

	return C
}

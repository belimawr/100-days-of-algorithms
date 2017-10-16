package main

import "errors"

func sumMatrix(a, b [][]int) ([][]int, error) {
	lA, okA := isSquareAndSize(a)
	lB, okB := isSquareAndSize(b)

	if !(okA && okB) {
		return [][]int{}, errors.New("The matrices are not square")
	}

	if lA != lB {
		return [][]int{}, errors.New("Matrices are not the same size")
	}

	n := lA

	result := make([][]int, n)

	for j := 0; j < n; j++ {
		result[j] = make([]int, n)
		for i := 0; i < n; i++ {
			result[j][i] = a[j][i] + b[j][i]
		}
	}

	return result, nil
}

func isSquareAndSize(m [][]int) (int, bool) {
	n := len(m)

	for _, v := range m {
		if len(v) != n {
			return 0, false
		}
	}

	return n, true
}

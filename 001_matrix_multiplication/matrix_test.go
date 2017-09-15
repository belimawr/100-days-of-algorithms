package main

import "testing"

func Test_printMatrix_1x1(t *testing.T) {
	A := [][]int{
		{2},
	}

	B := [][]int{
		{5},
	}

	C := [][]int{
		{10},
	}

	result := multiplyMatrix(A, B)

	for j := 0; j < 1; j++ {
		for i := 0; i < 1; i++ {
			if result[i][j] != C[i][j] {
				t.Errorf("Expecting [%d][%d] = %d, got %d", i, j, C[i][j], result[i][j])
			}
		}
	}
}

func Test_printMatrix_2x2_identity(t *testing.T) {
	A := [][]int{
		{42, 420},
		{420, 42},
	}

	B := [][]int{
		{1, 0},
		{0, 1},
	}

	C := [][]int{
		{42, 420},
		{420, 42},
	}

	result := multiplyMatrix(A, B)

	for j := 0; j < 2; j++ {
		for i := 0; i < 2; i++ {
			if result[i][j] != C[i][j] {
				t.Errorf("Expecting [%d][%d] = %d, got %d", i, j, C[i][j], result[i][j])
			}
		}
	}
}

func Test_printMatrix_2x2(t *testing.T) {
	A := [][]int{
		{1, 2},
		{3, 4},
	}

	B := [][]int{
		{2, 0},
		{1, 2},
	}

	C := [][]int{
		{4, 4},
		{10, 8},
	}

	result := multiplyMatrix(A, B)

	for j := 0; j < 2; j++ {
		for i := 0; i < 2; i++ {
			if result[i][j] != C[i][j] {
				t.Errorf("Expecting [%d][%d] = %d, got %d", i, j, C[i][j], result[i][j])
			}
		}
	}
}

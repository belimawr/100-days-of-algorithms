package main

import (
	"testing"
)

func Test_isSquareAndSize(t *testing.T) {
	expectedN := 2
	m := make([][]int, expectedN)

	for i := 0; i < 2; i++ {
		m[i] = make([]int, expectedN)
	}

	n, ok := isSquareAndSize(m)

	if !ok {
		t.Error("expecting true, got false")
	}

	if n != expectedN {
		t.Errorf("expecting %d, got %d", expectedN, n)
	}
}

func Test_isSquareAndSize_not_square(t *testing.T) {
	m := make([][]int, 2)

	for i := 0; i < 2; i++ {
		m[i] = make([]int, i+1)
	}

	n, ok := isSquareAndSize(m)

	if ok {
		t.Error("expecting false, got true")
	}

	if n != 0 {
		t.Errorf("expecting 0, got %d", n)
	}
}

func Test_sumMatrix(t *testing.T) {
	a := [][]int{
		[]int{1, 2},
		[]int{3, 4},
	}

	b := [][]int{
		[]int{6, 7},
		[]int{8, 9},
	}

	expected := [][]int{
		[]int{7, 9},
		[]int{11, 13},
	}

	m, err := sumMatrix(a, b)

	if err != nil {
		t.Errorf("did not expect an error: %q", err.Error())
	}

	n := len(a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if expected[i][j] != m[i][j] {
				t.Errorf("m[%d][%d] = %d, expecting %d",
					i, j,
					m[i][j],
					expected[i][j])
			}
		}
	}
}

func Test_sumMatrix_not_square(t *testing.T) {
	a := [][]int{
		[]int{1, 2},
		[]int{1, 2, 3},
	}

	b := [][]int{
		[]int{6, 7},
		[]int{8, 9},
	}

	m, err := sumMatrix(a, b)

	if err == nil {
		t.Error("expecting an error")
	}

	if len(m) != 0 {
		t.Error("returned matrix must be empty")
	}
}

func Test_sumMatrix_square_different_sizes(t *testing.T) {
	a := [][]int{
		[]int{1, 2},
		[]int{1, 2},
	}

	b := [][]int{
		[]int{6, 7, 0},
		[]int{8, 9, 0},
		[]int{8, 9, 0},
	}

	m, err := sumMatrix(a, b)

	if err == nil {
		t.Error("expecting an error")
	}

	if len(m) != 0 {
		t.Error("returned matrix must be empty")
	}
}

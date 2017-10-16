package main

import (
	"fmt"
)

func main() {
	var m, n int

	fmt.Scanf("%d, %d", &n, &m)
	a := readMatrix(n, m)

	fmt.Scanf("%d, %d", &n, &m)
	b := readMatrix(n, m)

	fmt.Printf("---\n\n")

	printMatrix(a)
	fmt.Printf("---\n")

	printMatrix(b)
	fmt.Printf("---\n")

	r, err := sumMatrix(a, b)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("\nResult:")
	printMatrix(r)
}

func printMatrix(m [][]int) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Print("\n")
	}
}

func readMatrix(n, m int) [][]int {
	r := make([][]int, n)

	for j := 0; j < n; j++ {
		r[j] = make([]int, m)
		for i := 0; i < m; i++ {
			var tmp int
			fmt.Scanf("%d", &tmp)
			r[j][i] = tmp
		}
	}

	return r
}

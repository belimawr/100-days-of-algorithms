package main

import "fmt"

func main() {
	var n int
	var tmp int64

	fmt.Scanf("%d", &n)

	v := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		v[i] = tmp
	}

	fmt.Println("Original:", v)
	bubbeSort(v)
	fmt.Println("Ordered: ", v)
}

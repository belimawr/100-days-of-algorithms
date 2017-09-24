package main

import (
	"fmt"
)

func main() {
	var size int
	var n int

	fmt.Print("Number to search: ")
	fmt.Scanf("%d", &n)

	fmt.Print("Size of the slice: ")
	fmt.Scanf("%d", &size)

	vec := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("vec[%d]: ", i)
		fmt.Scanf("%d", &(vec[i]))
	}

	idx := binarySearch(vec, n)

	if idx < 0 {
		fmt.Printf("Did not found %d on %#v\n", n, vec)
		return
	}

	fmt.Printf("Found %d on index: %d\n", n, idx)
}

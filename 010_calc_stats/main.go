package main

import (
	"fmt"
)

func main() {
	var size int

	fmt.Print("Size of the slice: ")
	fmt.Scanf("%d", &size)

	vec := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("vec[%d]: ", i)
		fmt.Scanf("%d", &(vec[i]))
	}

	fmt.Println("Min: ", minimun(vec))
	fmt.Println("Max: ", maximun(vec))
	fmt.Println("Count: ", count(vec))
	fmt.Println("Average: ", average(vec))
}

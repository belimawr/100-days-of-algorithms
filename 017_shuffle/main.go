package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var s, n, tmp int

	fmt.Print("Seed: ")
	fmt.Scanf("%d", &s)
	fmt.Print("Slice size: ")
	fmt.Scanf("%d", &n)

	rand.Seed(int64(s))

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("v[%d]: ", i)
		fmt.Scanf("%d", &tmp)
		v[i] = tmp
	}

	fmt.Println("Original:", v)
	shuffle(v)
	fmt.Println("Shufle: ", v)
}

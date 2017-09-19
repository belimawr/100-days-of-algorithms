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

	fmt.Println("Sequence:    ", v)
	next := nextPermutation([]int{4, 2, 5, 3, 1})
	fmt.Println("Permutation: ", next)

	// Uncomment the code below to see a sequence of permutations
	//x := []int{4, 3, 2, 1}
	//for i := 0; i < 25; i++ {
	//	x = nextPermutation(x)
	//	fmt.Println(x)
	//}
}

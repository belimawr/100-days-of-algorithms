package main

import (
	"flag"
	"fmt"
)

var A List
var B List
var C List

func main() {
	var n = flag.Uint("n", 5, "Number of disks")

	flag.Parse()

	for i := *n; i > 0; i-- {
		A.Push(i)
	}

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	Move(&A, &B, &C, *n)
}

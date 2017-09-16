package main

import (
	"flag"
	"fmt"
)

func main() {
	var n = flag.Uint64("n", 7, "Number to count the bits")

	flag.Parse()

	fmt.Printf("%d (0b%[1]b) has %d bits '1'\n", *n, countBits(*n))
}

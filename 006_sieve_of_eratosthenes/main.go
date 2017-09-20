package main

import (
	"flag"
	"fmt"
)

func main() {
	var n = flag.Int("n", 5, "Count primes until N")

	flag.Parse()

	fmt.Printf("The number of primes untin %d is %d\n", *n, sieve_of_eratosthenes(*n))
}

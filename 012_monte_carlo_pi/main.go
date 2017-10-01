package main

import (
	"flag"
	"fmt"
)

func main() {
	var n = flag.Int("n", 10000000, "Number of interactions")

	flag.Parse()

	fmt.Println(monteCarloPi(*n))
}

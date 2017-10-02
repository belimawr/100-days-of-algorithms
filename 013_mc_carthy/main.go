package main

import (
	"flag"
	"fmt"
)

func main() {
	var n = flag.Int("n", 42, "Number to calculate the McCarthy function")

	flag.Parse()

	mc := mcCarty(*n)

	fmt.Printf("McCarthy(%d) = %d\n", *n, mc)
}

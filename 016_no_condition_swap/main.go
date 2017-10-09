package main

import (
	"flag"
	"fmt"
)

func main() {
	var x = flag.Int("x", 42, "First number")
	var y = flag.Int("y", 1, "Second number")

	flag.Parse()

	swapIfgt(x, y)

	fmt.Printf("(x, y) => (%d, %d)\n", *x, *y)
}

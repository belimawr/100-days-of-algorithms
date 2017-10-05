package main

import (
	"flag"
	"fmt"
)

func main() {
	data := flag.String("d", "Hello, world!", "The string to work on")
	zeros := flag.Int("z", 4, "Number of leading zeros")

	flag.Parse()

	fmt.Printf("Got sha256 = %s\n", proofOfWork(*zeros, []byte(*data)))
}

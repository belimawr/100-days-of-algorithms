package main

import (
	"flag"
	"fmt"
)

func main() {
	var op1 = flag.String("op1", "011", "First operator")
	var op2 = flag.String("op2", "1011", "Second operator")

	flag.Parse()

	sum, err := binaryAdder(*op1, *op2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s + %s = %s\n",
		*op1,
		*op2,
		sum,
	)
}

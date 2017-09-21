package main

import (
	"flag"
	"fmt"
)

func main() {
	var expression = flag.String("exp", "40 2 +", "Expression to be calculated")

	flag.Parse()

	result, err := postfix_notation(*expression)

	if err != nil {
		fmt.Printf("There was an error: %q\n", err.Error())
		return
	}

	fmt.Printf("%s = %f\n", *expression, result)
}

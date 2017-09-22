package main

import (
	"fmt"
	"strings"
)

// Carry, op1, op2 -> result, carry
var transition = map[string][]string{
	"000": {"0", "0"},
	"001": {"1", "0"},
	"010": {"1", "0"},
	"011": {"0", "1"},

	"100": {"1", "0"},
	"101": {"0", "1"},
	"110": {"0", "1"},
	"111": {"1", "1"},
}

func binaryAdder(n1, n2 string) (string, error) {
	result := []string{}

	op1 := reversed(strings.Split(n1, ""))
	op2 := reversed(strings.Split(n2, ""))

	diff := len(op1) - len(op2)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			op2 = append(op2, "0")
		}
	} else {
		for i := 0; i < -diff; i++ {
			op1 = append(op1, "0")
		}
	}

	carry := "0"
	for i := range op1 {
		sum, ok := transition[carry+op1[i]+op2[i]]
		if !ok {
			return "", fmt.Errorf("one of the tokens: %q, %q is invalid",
				op1[i], op2[i])
		}

		result = append(result, sum[0])
		carry = sum[1]
	}

	if carry == "1" {
		result = append(result, carry)
	}

	return strings.Join(reversed(result), ""), nil
}

func reversed(s []string) []string {
	cp := append([]string(nil), s...)

	for left, right := 0, len(cp)-1; left < right; left, right = left+1, right-1 {
		cp[left], cp[right] = cp[right], cp[left]
	}

	return cp
}

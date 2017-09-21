package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []float64

func (l *Stack) Pop() float64 {
	last := len(*l)
	el := (*l)[last-1]

	*l = (*l)[:len(*l)-1]

	return el
}

func (l *Stack) Push(el float64) {
	*l = append(*l, el)
}

func (l *Stack) CanPop(amount int) bool {
	return len(*l) >= amount
}

var operations = map[string]func(float64, float64) float64{
	"+": func(a, b float64) float64 { return a + b },
	"-": func(a, b float64) float64 { return a - b },
	"*": func(a, b float64) float64 { return a * b },
	"/": func(a, b float64) float64 { return a / b },
}

func postfix_notation(expression string) (float64, error) {
	var n float64
	var err error

	split := strings.Split(expression, " ")
	stack := Stack{}

	for _, token := range split {
		operation, validOperation := operations[token]

		if validOperation {
			if !stack.CanPop(2) {
				return 0, fmt.Errorf("invalid expression, to few elements to peform '%s' operation", token)
			}

			op2 := stack.Pop()
			op1 := stack.Pop()
			n = operation(op1, op2)
		} else {
			n, err = strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
		}

		stack.Push(n)
	}

	return stack.Pop(), nil
}

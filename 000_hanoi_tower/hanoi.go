package main

import "fmt"

type Stack []uint

func (l *Stack) Pop() uint {
	last := len(*l)
	el := (*l)[last-1]

	*l = (*l)[:len(*l)-1]

	return el
}

func (l *Stack) Push(el uint) {
	*l = append(*l, el)
}

func Move(from, aux, to *Stack, h uint) {
	if h < 1 {
		return
	}

	Move(from, to, aux, h-1)

	p := from.Pop()
	to.Push(p)

	// TODO: Remove this global state dependency, probably using closure
	fmt.Println("#########################")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	Move(aux, from, to, h-1)
}

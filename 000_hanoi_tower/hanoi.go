package main

import "fmt"

type List struct {
	lst []uint
}

func (l *List) Pop() uint {
	el := l.lst[len(l.lst)-1]
	l.lst = l.lst[:len(l.lst)-1]

	return el
}

func (l *List) Push(el uint) {
	l.lst = append(l.lst, el)
}

func Move(from, aux, to *List, h uint) {
	if h < 1 {
		return
	}

	Move(from, to, aux, h-1)

	p := from.Pop()
	to.Push(p)

	fmt.Println("#########################")
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	Move(aux, from, to, h-1)
}

package main

import "testing"

func Test_swap_GT(t *testing.T) {
	x := 100
	y := 10

	swapIfgt(&x, &y)

	if y != 100 || x != 10 {
		t.Error("Did not swap")
	}
}

func Test_swap_LT(t *testing.T) {
	x := 10
	y := 100

	swapIfgt(&x, &y)

	if y != 100 || x != 10 {
		t.Error("Cannot swap")
	}
}

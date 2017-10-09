package main

func swapIfgt(x, y *int) {
	u, v := *x, *y
	s := (u - v) >> (31)

	*x = v*(1+s) - u*s
	*y = u*(1+s) - v*s
}

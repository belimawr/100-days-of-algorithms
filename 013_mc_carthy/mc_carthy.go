package main

func mcCarty(n int) int {
	if n > 100 {
		return n - 10
	}

	return mcCarty(mcCarty(n + 11))
}

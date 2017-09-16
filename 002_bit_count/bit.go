package main

func countBits(n uint64) uint8 {
	var c uint8

	for n > 0 {
		n &= n - 1
		c++
	}

	return c
}

package main

import (
	"math/rand"
)

func shuffle(v []int) {
	n := len(v)
	for i := 0; i < n; i++ {
		k := random(i, n)
		v[i], v[k] = v[k], v[i]
	}
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

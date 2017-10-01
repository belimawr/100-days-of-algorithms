package main

import (
	"math"
	"math/rand"
)

func monteCarloPi(n int) float64 {
	var count, countInside int

	for count = 0; count < n; count++ {
		x, y := rand.Float64(), rand.Float64()
		if math.Hypot(x, y) < float64(1) {
			countInside++
		}
	}

	return float64(4) * (float64(countInside) / float64(count))
}

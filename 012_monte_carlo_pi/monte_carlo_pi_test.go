package main

import (
	"math"
	"testing"
)

func Test_monteCarloPi(t *testing.T) {
	delta := 0.001
	pi := monteCarloPi(10000000)

	if math.Abs(pi-math.Pi) > delta {
		t.Errorf("monteCarloPi - math.Pi: %f - %f = %f", pi, math.Pi, math.Abs(pi-math.Pi))
	}
}

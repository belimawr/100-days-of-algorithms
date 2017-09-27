package main

import (
	"math"
	"testing"
)

var vec = []int{6, 9, 15, -2, 92, 11}

func Test_minimun(t *testing.T) {
	min := minimun(vec)

	if min != -2 {
		t.Errorf("Expecting -2, got %d", min)
	}
}

func Test_minimiun_nil(t *testing.T) {
	min := minimun(nil)

	if min != math.MinInt32 {
		t.Errorf("Expecting %d, got %d", math.MinInt32, min)
	}
}

func Test_maximun(t *testing.T) {
	max := maximun(vec)

	if max != 92 {
		t.Errorf("Expecing 92, got %d", max)
	}
}

func Test_maximun_nil(t *testing.T) {
	max := maximun(nil)

	if max != math.MaxInt32 {
		t.Errorf("Expecing %d, got %d", math.MaxInt32, max)
	}
}

func Test_count(t *testing.T) {
	expected := len(vec)
	c := count(vec)

	if c != expected {
		t.Errorf("Expecting %d, got %d", expected, c)
	}
}

func Test_count_nil(t *testing.T) {
	expected := len([]int(nil))
	c := count(nil)

	if c != expected {
		t.Errorf("Expecting %d, got %d", expected, c)
	}
}

func Test_average(t *testing.T) {
	expected := 21.833333
	avg := average(vec)
	delta := 0.000001

	if math.Abs(avg-expected) > delta {
		t.Errorf("Expecting |%.6f - %.6f| < %.6f, got %.6f",
			expected,
			avg,
			delta,
			math.Abs(average(vec)-expected))
	}
}

func Test_average_nil(t *testing.T) {
	avg := average(nil)

	if !math.IsNaN(avg) {
		t.Errorf("Expecting NaN, got %.6f", avg)
	}
}

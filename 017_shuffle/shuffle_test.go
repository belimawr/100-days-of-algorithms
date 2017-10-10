package main

import (
	"math/rand"
	"testing"
)

func Test_random(t *testing.T) {
	rand.Seed(1)
	n := 10

	expectedResults := []int{1, 7, 9, 3, 5, 8, 7, 9, 8, 9}

	for i := 0; i < n; i++ {
		r := random(i, n)
		if r != expectedResults[i] {
			t.Errorf("Expecting %d got %d", expectedResults[i], r)
		}
	}
}

func Test_shuffle(t *testing.T) {
	v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{1, 7, 9, 3, 5, 8, 0, 2, 4, 6}

	rand.Seed(1)
	shuffle(v)

	for i := 0; i < len(v); i++ {
		if v[i] != expected[i] {
			t.Errorf("expecting %d, got %d", expected[i], v[i])
		}
	}
}

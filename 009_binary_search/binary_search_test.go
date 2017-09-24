package main

import "testing"

func Test_binarySearch(t *testing.T) {
	vec := []int{1, 2, 3}
	n := 2
	expected := 1

	idx := binarySearch(vec, n)

	if idx != expected {
		t.Errorf("Expecting idx = %d, got %d", idx, expected)
	}
}

func Test_binarySearch_last(t *testing.T) {
	vec := []int{1, 2, 3}
	n := 3
	expected := 2

	idx := binarySearch(vec, n)

	if idx != expected {
		t.Errorf("Expecting idx = %d, got %d", idx, expected)
	}
}

func Test_binarySearch_first(t *testing.T) {
	vec := []int{1, 2, 3}
	n := 1
	expected := 0

	idx := binarySearch(vec, n)

	if idx != expected {
		t.Errorf("Expecting idx = %d, got %d", idx, expected)
	}
}

func Test_binarySearch_not_found_left(t *testing.T) {
	vec := []int{1, 2, 3}
	n := -42
	expected := -1

	idx := binarySearch(vec, n)

	if idx != expected {
		t.Errorf("Expecting idx = %d, got %d", idx, expected)
	}
}

func Test_binarySearch_not_found_right(t *testing.T) {
	vec := []int{1, 2, 3}
	n := 42
	expected := -1

	idx := binarySearch(vec, n)

	if idx != expected {
		t.Errorf("Expecting idx = %d, got %d", idx, expected)
	}
}

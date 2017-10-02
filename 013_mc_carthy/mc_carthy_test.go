package main

import "testing"

func Test_mcCarty_lt_100(t *testing.T) {
	expected := 91

	n := mcCarty(42)

	if n != expected {
		t.Errorf("Expecting %d got %d", expected, n)
	}
}

func Test_mcCarty_gte_100(t *testing.T) {
	expected := 93

	n := mcCarty(103)

	if n != expected {
		t.Errorf("Expecting %d got %d", expected, n)
	}
}

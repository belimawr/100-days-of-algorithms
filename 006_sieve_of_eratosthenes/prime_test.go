package main

import "testing"

func Test_sieve_of_eratosthenes_until_2(t *testing.T) {
	expected := 1
	got := sieve_of_eratosthenes(2)

	if got != expected {
		t.Errorf("Expecting %d, got %d", expected, got)
	}
}

func Test_sieve_of_eratosthenes_until_5(t *testing.T) {
	expected := 3
	got := sieve_of_eratosthenes(5)

	if got != expected {
		t.Errorf("Expecting %d, got %d", expected, got)
	}
}

func Test_sieve_of_eratosthenes_until_10(t *testing.T) {
	expected := 4
	got := sieve_of_eratosthenes(10)

	if got != expected {
		t.Errorf("Expecting %d, got %d", expected, got)
	}
}

func Test_sieve_of_eratosthenes_until_1000(t *testing.T) {
	expected := 168
	got := sieve_of_eratosthenes(1000)

	if got != expected {
		t.Errorf("Expecting %d, got %d", expected, got)
	}
}

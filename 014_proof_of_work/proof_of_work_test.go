package main

import "testing"

func Test_countZeros(t *testing.T) {
	tests := []struct {
		expected int
		str      string
	}{
		{
			expected: 0,
			str:      "42",
		},

		{
			expected: 10,
			str:      "0000000000 123123",
		},

		{
			expected: 1,
			str:      "0zero0",
		},
	}

	for _, test := range tests {
		got := countZeros(test.str)

		if got != test.expected {
			t.Errorf("%s -> got %d expecting %d", test.str, got, test.expected)
		}
	}
}

func Test_proofOfWork(t *testing.T) {
	expected := 4

	s := proofOfWork(expected, []byte("Hello, world!42"))

	count := 0
	for i := 0; i < len(s) && s[i] == '0'; i++ {
		count++
	}

	if count != expected {
		t.Errorf("Got %d zeros, expecting %d", count, expected)
	}
}

package main

import "testing"

func Test_countBit_0b1(t *testing.T) {
	// 1 = 0b1
	n := countBits(1)

	if n != 1 {
		t.Errorf("Got %d, expecing %d", n, 1)
	}
}

func Test_countBit_0b0(t *testing.T) {
	// 0 = 0b1=0
	n := countBits(0)

	if n != 0 {
		t.Errorf("Got %d, expecing %d", n, 0)
	}
}

func Test_countBit_0b01100101100000(t *testing.T) {
	// 6496 = 0b01100101100000
	n := countBits(6496)

	if n != 5 {
		t.Errorf("Got %d, expecing %d", n, 5)
	}
}

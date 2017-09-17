package main

import (
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	v := []int64{4, 3, 2, 1, 0}

	bubbeSort(v)

	for i := 0; i < len(v); i++ {
		if v[i] != int64(i) {
			t.Errorf("Expecting v[%d] = %[1]d, got %d", i, v[i])
		}
	}
}

func Test_bubbleSort_nil_should_not_panic(t *testing.T) {
	bubbeSort(nil)
}

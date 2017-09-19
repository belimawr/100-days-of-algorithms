package main

import (
	"reflect"
	"testing"
)

func Test_reversed(t *testing.T) {
	lst := []int{0, 1, 2, 3}

	rev := reversed(lst[:])

	for i := len(lst)/2 - 1; i >= 0; i-- {
		tmp := len(lst) - 1 - i
		if lst[i] != rev[tmp] {
			t.Errorf("lst, rev: %d != %d", lst[i], rev[tmp])
		}
	}
}

func Test_nextPermutation(t *testing.T) {
	perm := nextPermutation([]int{4, 2, 5, 3, 1})

	if !reflect.DeepEqual(perm, []int{4, 3, 1, 2, 5}) {
		t.Error("Wrong permutation")
	}
}

func Test_nextPermutation_one_element(t *testing.T) {
	perm := nextPermutation([]int{4})

	if !reflect.DeepEqual(perm, []int{4}) {
		t.Error("Wrong permutation")
	}
}

func Test_nextPermutation_two_elements(t *testing.T) {
	perm := nextPermutation([]int{4, 2})

	if !reflect.DeepEqual(perm, []int{2, 4}) {
		t.Error("Wrong permutation")
	}
}

func Test_nextPermutation_nil(t *testing.T) {
	perm := nextPermutation(nil)

	if !reflect.DeepEqual(perm, []int(nil)) {
		t.Error("Wrong permutation")
	}
}

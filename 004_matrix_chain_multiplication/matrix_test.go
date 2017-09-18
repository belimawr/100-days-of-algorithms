package main

import "testing"

func Test_matrixChainOrder(t *testing.T) {
	dims := []int64{40, 20, 30, 10, 30}
	expectedOrder := "((A(BC))D)"
	expectedCount := int64(26000)

	order, cost := matrixChainOrder(dims)

	if cost != expectedCount {
		t.Errorf("Expecting %d, got %d", expectedCount, cost)
	}

	if order != expectedOrder {
		t.Errorf("Got %q, expecting %q", order, expectedOrder)
	}
}

func Test_matrixChainOrder_two_matrices(t *testing.T) {
	dims := []int64{40, 20, 30}
	expectedOrder := "(AB)"
	expectedCount := int64(24000)

	order, cost := matrixChainOrder(dims)

	if cost != expectedCount {
		t.Errorf("Expecting %d, got %d", expectedCount, cost)
	}

	if order != expectedOrder {
		t.Errorf("Got %q, expecting %q", order, expectedOrder)
	}
}

func Test_matrixChainOrder_too_small(t *testing.T) {
	dims := []int64{40, 20}
	expectedOrder := ""
	expectedCount := int64(0)

	order, cost := matrixChainOrder(dims)

	if cost != expectedCount {
		t.Errorf("Expecting %d, got %d", expectedCount, cost)
	}

	if order != expectedOrder {
		t.Errorf("Got %q, expecting %q", order, expectedOrder)
	}
}

func Test_matrixChainOrder_nil(t *testing.T) {
	expectedOrder := ""
	expectedCount := int64(0)

	order, cost := matrixChainOrder(nil)

	if cost != expectedCount {
		t.Errorf("Expecting %d, got %d", expectedCount, cost)
	}

	if order != expectedOrder {
		t.Errorf("Got %q, expecting %q", order, expectedOrder)
	}
}

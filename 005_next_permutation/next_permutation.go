package main

func nextPermutation(values []int) []int {
	last := len(values) - 1
	pivot := -1

	// Find the pivot
	for i := last - 1; i >= 0; i-- {
		if values[i] < values[i+1] {
			pivot = i
			break
		}
	}

	// If it wasn't found, return the reversed list
	if pivot == -1 {
		return reversed(values)
	}

	// Find the first element bigger than the pivot in the sub-sequence (from end to beginning)
	for i := last; i >= pivot; i-- {
		if values[pivot] < values[i] {
			// Swap pivot and the element
			values[pivot], values[i] = values[i], values[pivot]
			// Reverse the sub-sequence
			values = append(values[:pivot+1], reversed(values[pivot+1:])...)
		}
	}

	return values
}

func reversed(s []int) []int {
	cp := append([]int(nil), s...)

	for left, right := 0, len(cp)-1; left < right; left, right = left+1, right-1 {
		cp[left], cp[right] = cp[right], cp[left]
	}

	return cp
}

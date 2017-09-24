package main

import "math"

func binarySearch(vec []int, n int) int {
	var left, right float64
	left, right = 0, float64(len(vec)-1)

	var middle int

	for left <= right {
		middle = int(math.Floor((left + right) / 2))

		if vec[middle] == n {
			return middle
		} else if vec[middle] > n {
			right = float64(middle - 1)
		} else {
			left = float64(middle + 1)

		}
	}

	return -1
}

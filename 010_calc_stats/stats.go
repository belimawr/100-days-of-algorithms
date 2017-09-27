package main

import "math"

func minimun(vec []int) int {
	if len(vec) == 0 {
		return math.MinInt32
	}

	min := vec[0]

	for i := 1; i < len(vec); i++ {
		if min > vec[i] {
			min = vec[i]
		}
	}

	return min
}

func maximun(vec []int) int {
	if len(vec) == 0 {
		return math.MaxInt32
	}

	max := vec[0]

	for i := 1; i < len(vec); i++ {
		if max < vec[i] {
			max = vec[i]
		}
	}

	return max
}

func count(vec []int) int {
	c := 0

	for range vec {
		c++
	}

	return c
}

func average(vec []int) float64 {
	var avg float64

	for _, v := range vec {
		avg += float64(v)
	}

	return avg / float64(len(vec))
}

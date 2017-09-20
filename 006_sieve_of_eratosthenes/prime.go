package main

func sieve_of_eratosthenes(n int) int {
	numbers := make(map[int]bool)

	for i := 2; i <= n; i++ {
		numbers[i] = true
	}

	for prime := 2; prime <= n; prime++ {
		for i := prime * prime; i <= n; i += prime {
			delete(numbers, i)
		}
	}

	return len(numbers)
}

package main

func bubbeSort(v []int64) {
	for i := 0; i < len(v); i++ {
		for j := i; j < len(v); j++ {
			if v[i] > v[j] {
				tmp := v[i]
				v[i] = v[j]
				v[j] = tmp
			}
		}
	}
}

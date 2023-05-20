package main

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	increases := true
	decreases := true

	for i := 1; i < len(array); i++ {
		if !(array[i-1] <= array[i]) {
			increases = false
		}
		if !(array[i-1] >= array[i]) {
			decreases = false
		}
	}
	return increases || decreases
}

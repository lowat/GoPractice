package main

import "math"

func FindThreeLargestNumbers(array []int) []int {
	largestValues := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, value := range array {
		placeValue(largestValues, value)
	}
	return largestValues
}

func placeValue(largestValues []int, value int) {
	for index := len(largestValues) - 1; index >= 0; index-- {
		if value > largestValues[index] {
			value = replace_value_at_index_with_new_value_and_return(largestValues, value, index)
		}
	}
}

func replace_value_at_index_with_new_value_and_return(largestValues []int, newValue int, index int) int {
	displacedValue := largestValues[index]
	largestValues[index] = newValue
	return displacedValue
}

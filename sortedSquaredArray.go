package main

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	sortedSquares := make([]int, len(array))

	var smallerValueIdx int = 0
	var largerValueIdx int = len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallerValue := array[smallerValueIdx]
		largerValue := array[largerValueIdx]

		if abs(smallerValue) > abs(largerValue) {
			sortedSquares[idx] = smallerValue * smallerValue
			smallerValueIdx += 1
		} else {
			sortedSquares[idx] = largerValue * largerValue
			largerValueIdx -= 1
		}
	}

	return sortedSquares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

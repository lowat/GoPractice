package main

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	arrPointer := 0
	seqPointer := 0

	for seqPointer < len(sequence) {
		if arrPointer >= len(array) {
			return false
		}
		if array[arrPointer] == sequence[seqPointer] {
			seqPointer++
		}
		arrPointer++
	}
	return seqPointer == len(sequence)
}

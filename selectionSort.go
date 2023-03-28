package main

func SelectionSort(array []int) []int {
	currentIdx := 0
	for currentIdx < len(array)-1 {
		smallestIdx := currentIdx
		for i := smallestIdx + 1; i < len(array); i++ {
			if array[i] < array[smallestIdx] {
				smallestIdx = i
			}
		}
		swap(array, smallestIdx, currentIdx)
		currentIdx++
	}
	return array
}

func swap(array []int, indexOne int, indexTwo int) {
	temp := array[indexTwo]
	array[indexTwo] = array[indexOne]
	array[indexOne] = temp
}

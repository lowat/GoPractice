package main

func BubbleSort(array []int) []int {
	isSorted := false
	endPointer := len(array) - 1
	for !isSorted {
		isSorted = true
		for i := 0; i < endPointer; i++ {
			if array[i] > array[i+1] {
				swap(array, i, i+1)
				isSorted = false
			}
		}
		endPointer--
	}
	return array
}

func swap(array []int, indexOne int, indexTwo int) {
	temp := array[indexOne]
	array[indexOne] = array[indexTwo]
	array[indexTwo] = temp
}

package main

func MoveElementToEnd(array []int, toMove int) []int {
	left := 0
	right := len(array) - 1
	for left < right {
		leftValue := array[left]
		rightValue := array[right]
		if leftValue == toMove && rightValue != toMove {
			swap(left, right, array)
		}
		if leftValue != toMove {
			left++
		}
		if rightValue == toMove {
			right--
		}
	}
	return array
}

func swap(idxOne int, idxTwo int, array []int) {
	temp := array[idxOne]
	array[idxOne] = array[idxTwo]
	array[idxTwo] = temp
}

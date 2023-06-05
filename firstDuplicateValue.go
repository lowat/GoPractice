package main

func FirstDuplicateValue(array []int) int {
	for _, value := range array {
		absValue := abs(value)
		if array[absValue-1] < 0 {
			return absValue
		}
		array[absValue-1] *= -1
	}
	return -1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

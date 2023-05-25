package main

func ArrayOfProducts(array []int) []int {
	result := []int{}
	leftProduct := 1
	for i := 0; i < len(array); i++ {
		result = append(result, leftProduct)
		leftProduct *= array[i]
	}
	rightProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct *= array[i]
	}
	return result
}

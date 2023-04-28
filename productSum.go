package main

type SpecialArray []interface{}

func ProductSum(array SpecialArray) int {
	return helper(array, 1)
}

func helper(array SpecialArray, depth int) int {
	sum := 0
	for _, el := range array {
		if cast, ok := el.(SpecialArray); ok {
			sum += helper(cast, depth+1)
		} else if cast, ok := el.(int); ok {
			sum += cast
		}
	}
	return sum * depth
}

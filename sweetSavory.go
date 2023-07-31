package main

import "sort"

func SweetAndSavory(dishes []int, target int) []int {
	sort.Ints(dishes)
	closestDiff := int(^uint(0) >> 1)
	left := 0
	right := len(dishes) - 1
	result := []int{0, 0}
	for left < right && dishes[left] < 0 && dishes[right] > 0 {
		currentPairValue := dishes[left] + dishes[right]
		if currentPairValue > target {
			right--
			continue
		}
		currentDiff := target - currentPairValue
		if currentDiff < closestDiff {
			result[0] = dishes[left]
			result[1] = dishes[right]
			closestDiff = currentDiff
			if currentPairValue == target {
				return result
			}
		}
		left++
	}
	return result

	return []int{}
}

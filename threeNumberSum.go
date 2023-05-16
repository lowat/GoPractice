package main

import (
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	return findTriplets(array, target)
}

func findTriplets(array []int, target int) [][]int {
	triplets := [][]int{}
	for i := 0; i < len(array)-2; i++ {
		left := i + 1
		right := len(array) - 1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum == target {
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left++
				right--
			} else if currentSum > target {
				right--
			} else {
				left++
			}
		}
	}
	return triplets
}

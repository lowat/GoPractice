package main

import (
	"sort"
)

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	var currentMaxChange = 0
	for _, coin := range coins {
		if coin > currentMaxChange+1 {
			return currentMaxChange + 1
		}
		currentMaxChange += coin
	}
	return currentMaxChange + 1
}

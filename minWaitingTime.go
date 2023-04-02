package main

import (
	"sort"
)

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)
	timeWaited := 0
	queriesRan := 0
	for i := 0; i < len(queries)-1; i++ {
		queriesRan += queries[i]
		timeWaited += queriesRan
	}
	return timeWaited
}

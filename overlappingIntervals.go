package main

import (
	"sort"
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	sorted := make([][]int, len(intervals))
	copy(sorted, intervals)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})

	result := make([][]int, 0)
	ongoingRange := sorted[0]
	for i := 1; i < len(sorted); i++ {
		if ongoingRange[1] >= sorted[i][0] {
			ongoingRange[1] = max(ongoingRange[1], sorted[i][1])
		} else {
			result = append(result, ongoingRange)
			ongoingRange = sorted[i]
		}
	}
	result = append(result, ongoingRange)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)
	return calcTandemSpeed(redShirtSpeeds, blueShirtSpeeds, fastest)
}

func calcTandemSpeed(tandemPoolA []int, tandemPoolB []int, fastest bool) int {
	indexA := 0
	indexB := 0
	deltaA := 1
	deltaB := 1
	if fastest {
		indexB = len(tandemPoolB) - 1
		deltaB = -1
	}
	tandemSpeed := 0
	for areIndicesValid(indexA, indexB, len(tandemPoolA), len(tandemPoolB)) {
		tandemSpeed += max(tandemPoolA[indexA], tandemPoolB[indexB])
		indexA += deltaA
		indexB += deltaB
	}
	return tandemSpeed
}

func areIndicesValid(indexA int, indexB int, listA_length int, listB_length int) bool {
	return indexA >= 0 && indexA < listA_length && indexB >= 0 && indexB < listB_length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

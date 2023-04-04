package main

import (
	"sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)
	if redShirtHeights[0] > blueShirtHeights[0] {
		return canTakePhoto(redShirtHeights, blueShirtHeights)
	}
	return canTakePhoto(blueShirtHeights, redShirtHeights)
}

func canTakePhoto(topRow []int, bottomRow []int) bool {
	for i, _ := range topRow {
		if bottomRow[i] >= topRow[i] {
			return false
		}
	}
	return true
}

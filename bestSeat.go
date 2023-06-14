package main

func BestSeat(seats []int) int {
	left := 0
	bestSeat := -1
	maxSpace := 0

	for left < len(seats) {
		right := left + 1

		for right < len(seats) && seats[right] == 0 {
			right++
		}

		currentSpace := right - left - 1

		if currentSpace > maxSpace {
			maxSpace = currentSpace
			bestSeat = (right + left) / 2
		}

		left = right
	}
	return bestSeat
}

package main

func MissingNumbers(nums []int) []int {
	targetLine := getTargetLine(nums)
	targetHalfSums := getTargetHalfSums(len(nums), targetLine)
	actualHalfSums := getActualHalfSums(nums, targetLine)
	firstHalfMissing := targetHalfSums[0] - actualHalfSums[0]
	secondHalfMissing := targetHalfSums[1] - actualHalfSums[1]
	return []int{firstHalfMissing, secondHalfMissing}
}

func getTargetLine(nums []int) int {
	total := 0
	for i := 1; i <= len(nums)+2; i++ {
		total += i
	}
	for _, num := range nums {
		total -= num
	}
	return total / 2
}

func getTargetHalfSums(numsLength, targetLine int) [2]int {
	targetFirstHalfSum := 0
	targetSecondHalfSum := 0
	for i := 1; i <= numsLength+2; i++ {
		if i <= targetLine {
			targetFirstHalfSum += i
		} else {
			targetSecondHalfSum += i
		}
	}
	return [2]int{targetFirstHalfSum, targetSecondHalfSum}
}

func getActualHalfSums(nums []int, targetLine int) [2]int {
	actualFirstHalfSum := 0
	actualSecondHalfSum := 0
	for _, num := range nums {
		if num <= targetLine {
			actualFirstHalfSum += num
		} else {
			actualSecondHalfSum += num
		}
	}
	return [2]int{actualFirstHalfSum, actualSecondHalfSum}
}

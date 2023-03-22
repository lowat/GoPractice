package main

func FirstNonRepeatingCharacter(str string) int {
	charFrequencies := map[rune]int{}

	for _, char := range str {
		charFrequencies[char] = charFrequencies[char] + 1
	}

	for i, char := range str {
		if charFrequencies[char] == 1 {
			return i
		}
	}

	return -1
}

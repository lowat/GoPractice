package main

func GenerateDocument(characters string, document string) bool {
	charMap := map[rune]int{}

	for _, char := range characters {
		charMap[char] = charMap[char] + 1
	}

	for _, char := range document {
		if charMap[char] == 0 {
			return false
		}

		charMap[char] = charMap[char] - 1
	}

	return true
}

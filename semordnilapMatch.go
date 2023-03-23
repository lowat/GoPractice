package main

func Semordnilap(words []string) [][]string {
	wordSet := map[string]bool{}
	for _, word := range words {
		wordSet[word] = true
	}

	semordnilapPairs := [][]string{}
	for _, word := range words {
		reverse := reverse(word)
		if _, reverseInSet := wordSet[reverse]; reverseInSet && reverse != word {
			semordnilapPairs = append(semordnilapPairs, []string{word, reverse})
			delete(wordSet, word)
			delete(wordSet, reverse)
		}
	}
	return semordnilapPairs
}

func reverse(word string) string {
	reverse := []byte{}
	for i := len(word) - 1; i >= 0; i-- {
		reverse = append(reverse, word[i])
	}
	return string(reverse)
}

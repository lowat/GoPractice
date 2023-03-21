import (
	"strconv"
)

func RunLengthEncoding(str string) string {
	encodedChars := []byte{}
	currentCharCount := 1

	for i := 1; i < len(str); i++ {
		currentChar := str[i]
		prevChar := str[i-1]
		if currentChar != prevChar || currentCharCount >= 9 {
			encodedChars = append(encodedChars, strconv.Itoa(currentCharCount)[0])
			encodedChars = append(encodedChars, prevChar)
			currentCharCount = 0
		}
		currentCharCount++
	}
	encodedChars = append(encodedChars, strconv.Itoa(currentCharCount)[0])
	encodedChars = append(encodedChars, str[len(str)-1])
	return string(encodedChars)

}
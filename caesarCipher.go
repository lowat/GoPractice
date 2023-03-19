import (
	"strings"
)

func CaesarCipherEncryptor(str string, key int) string {
	ALPHABET := "abcdefghijklmnopqrstuvwxyz"
	encryptedString := []rune(str)
	for i, char := range encryptedString {
		index := strings.Index(ALPHABET, string(char))
		if index == -1 {
			return ""
		}
		newIndex := (index + key) % 26
		encryptedString[i] = rune(ALPHABET[newIndex])
	}
	return string(encryptedString)
}
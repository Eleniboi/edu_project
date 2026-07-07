package transformer

import (
	"strings"
	"unicode"
)

func upperPrevious(result []string, n int) {
	if n > len(result) {
		n = len(result)
	}
	for i := len(result) - n; i < len(result); i++ {
		result[i] = strings.ToUpper(result[i])
	}
}
func lowerPrevious(result []string, n int) {
	if n > len(result) {
		n = len(result)
	}
	for i := len(result) - n; i < len(result); i++ {
		result[i] = strings.ToLower(result[i])
	}
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(strings.ToLower(word))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
func capPrevious(result []string, n int) {
	if n > len(result) {
		n = len(result)
	}
	for i := len(result) - n; i < len(result); i++ {
		result[i] = capitalize(result[i])
	}
}

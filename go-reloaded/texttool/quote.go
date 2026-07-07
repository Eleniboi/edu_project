package texttool

import (
	"strings"
)

func Quote(s string) string {

	text := strings.Split(s, "'")

	for i := 0; i < len(text); i++ {

		if i%2 == 1 {
			text[i] = strings.TrimSpace(text[i])
		}
	}
	word := strings.Join(text, "'")

	result := strings.Split(word, `"`)

	for i := 0; i < len(result); i++ {

		if i%2 == 1 {
			result[i] = strings.TrimSpace(result[i])
		}
	}
	return strings.Join(result, `"`)
}

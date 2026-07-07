package tokenizer

import "unicode"

func Tokenizer(text string) []string {
	var tokens []string
	var current string
	for i := 0; i < len(text); {
		ch := rune(text[i])

		if unicode.IsLetter(ch) {
			current = ""
			for i < len(text) && unicode.IsLetter(rune(text[i])) {
				current += string(text[i])
				i++
			}
			tokens = append(tokens, current)

		} else if unicode.IsDigit(ch) {
			current = ""
			for i < len(text) && unicode.IsDigit(rune(text[i])) {
				current += string(text[i])
				i++
			}
			tokens = append(tokens, current)

		} else if unicode.IsSpace(ch) {
			i++

		} else if ch == '(' {
			current = ""
			for i < len(text) && text[i] != ')' {
				current += string(text[i])
				i++
			}
			if i < len(text) {
				current += string(text[i])
				i++
			}
			tokens = append(tokens, current)
		} else {
			current = ""
			for i < len(text) &&
				!unicode.IsLetter(rune(text[i])) &&
				!unicode.IsDigit(rune(text[i])) &&
				!unicode.IsSpace(rune(text[i])) {

				current += string(text[i])
				i++
			}
			tokens = append(tokens, current)
		}
	}
	return tokens
}

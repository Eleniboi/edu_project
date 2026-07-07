package transformer

import "unicode"

func FixPunctuation(tokens []string) []string {

	for i := 0; i < len(tokens); i++ {
		if isPunctuationToken(tokens[i]) && i > 0 {

			tokens[i-1] = tokens[i-1] + tokens[i]

			tokens = append(tokens[:i], tokens[i+1:]...)
			i++
		}
	}
	return tokens
}

func isPunctuationToken(token string) bool {

	if token == "" {
		return false
	}

	for _, r := range token {
		if !unicode.IsPunct(r) {
			return false
		}
	}
	return true
}

package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {

	var b strings.Builder

	if input == "" {
		return ""
	}

	siceINput := SplitInput(input)

	isOnlynewline := true

	for _, word := range siceINput {
		if word != "" {
			isOnlynewline = false
		}
	}

	if isOnlynewline {
		for i := 0; i < len(siceINput)-1; i++ {
			b.WriteString("\n")
		}
		return b.String()
	}
	for _, word := range siceINput {

		if word == "" {
			// if len(siceINput) == 2 && siceINput[len(siceINput)-1] == "" {

			// 	asciiTable := RenderLine(word, banner)

			// 	for i := 0; i < len(asciiTable); i++ {
			// 		b.WriteString(asciiTable[i] + "\n")

			// 	}
			// 	return b.String()
			// }
			b.WriteString("\n")
			continue
		}

		asciiTable := RenderLine(word, banner)

		for i := 0; i < len(asciiTable); i++ {
			b.WriteString(asciiTable[i])
			b.WriteString("\n")
		}
	}
	return b.String()
}

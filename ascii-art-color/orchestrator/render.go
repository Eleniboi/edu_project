package orchestrator

import (
	"strings"
)

func renderInput(input, substr, colorCode string, banner map[rune][]string) []string {

	var result []string
	var build strings.Builder

	indexes := findIndexes(substr, input)

	for i := 0; i < 8; i++ {

		for indx, ch := range input {

			asciiPart := banner[ch][i]

			if substr == "" {

				asciiPart = colorCode + asciiPart + RESET
			}
			for _, start := range indexes {

				if start <= indx && indx < start+len(substr) {

					asciiPart = colorCode + asciiPart + RESET

				}
			}
			build.WriteString(asciiPart)
		}
		result = append(result, build.String())
		build.Reset()
	}
	return result
}

func findIndexes(substr, fullstr string) []int {

	var indexes []int

	if substr == "" {
		return indexes
	}

	offset := 0
	remaining := fullstr

	for {

		currentIndex := strings.Index(remaining, substr)

		if currentIndex == -1 {
			break
		}

		realIndex := offset + currentIndex
		indexes = append(indexes, realIndex)

		move := currentIndex + len(substr)

		offset += move
		remaining = remaining[move:]
	}
	return indexes
}

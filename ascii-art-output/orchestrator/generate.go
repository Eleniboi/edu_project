package orchestrator

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {

	var build strings.Builder

	if input == "" {
		return ""
	}

	slitInput := SplitInput(input)
	allempty := true

	for _, word := range slitInput {
		if word != "" {
			allempty = false
		}
	}

	if allempty {

		for i := 0; i < len(slitInput)-1; i++ {
			build.WriteString("\n")
		}
		return build.String()
	}

	for _, word := range slitInput {

		if word == "" {
			build.WriteString("\n")
			continue
		}

		result := RenderLine(word, banner)

		for i := 0; i < len(result); i++ {

			build.WriteString(result[i])
			build.WriteString("\n")
		}
	}
	return build.String()

}

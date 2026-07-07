package orchestrator

import "strings"

func GenerateArt(input, substr, colorCode string, banner map[rune][]string) string {

	var build strings.Builder

	if input == "" {
		return ""
	}

	slitinput := SplitInput(input)
	allempty := true

	for _, word := range slitinput {
		if word != "" {
			allempty = false
		}
	}

	if allempty {

		for i := 0; i < len(slitinput)-1; i++ {
			build.WriteString("\n")
		}
		return build.String()
	}

	for _, word := range slitinput {

		if word == "" {
			build.WriteString("\n")
			continue
		}

		result := renderInput(word, substr, colorCode, banner)

		for i := 0; i < len(result); i++ {

			build.WriteString(result[i])
			build.WriteString("\n")
		}
	}
	return build.String()

}

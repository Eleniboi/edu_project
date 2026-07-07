package orchestrator

import "strings"

func RenderLine(input string, banner map[rune][]string) []string {

	var result []string

	var build strings.Builder

	for i := 0; i < 8; i++ {

		for _, char := range input {
			build.WriteString(banner[char][i])
		}
		result = append(result, build.String())
		build.Reset()
	}
	return result
}

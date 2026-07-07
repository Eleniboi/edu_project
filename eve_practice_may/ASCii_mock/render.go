package main

import (
	"strings"
)

func RenderLine(input string, banner map[rune][]string) []string {

	var result []string
	var build strings.Builder

	for i := 0; i < 8; i++ {

		for _, r := range input {

			build.WriteString(banner[r][i])
		}
		result = append(result, build.String())
		build.Reset()
	}

	return result
}

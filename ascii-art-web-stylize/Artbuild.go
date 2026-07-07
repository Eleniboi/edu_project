package main

import (
	"os"
	"strings"
	"errors"
)

func Artbuilder(text, banner string) (string, error) {

	text = strings.ReplaceAll(text, "\r\n", "\n")

	file, err := os.ReadFile("banners/"+banner)

	if err != nil {
		return "", errors.New("error reading file")
	}

	fileline := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	fileline = fileline[1:]

	var build strings.Builder

	splitIlnput := strings.Split(text, "\n")

	for _, word := range splitIlnput {

		for row := 0; row < 8; row++ {

			for _, ch := range word {

				start := int(ch-32) * 9

				end := start + 8

				build.WriteString(fileline[start:end][row])
			}
			build.WriteString("\n")
		}
	}
	return build.String(), nil
}

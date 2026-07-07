package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func LoadBanner(bannerfile string) (map[rune][]string, error) {

	file, err := os.ReadFile(bannerfile)

	if err != nil {
		return nil, errors.New("Error reading file")
	}

	if len(file) == 0 {
		return nil, fmt.Errorf("Banner file is empty")
	}

	if len(file) < 855{
		return  nil, errors.New("Incomplete file content")
	}
	fileLines := strings.Split(string(file), "\n")

	fileLines = fileLines[1:]

	var banner = make(map[rune][]string)
	for i := ' '; i <= '~'; i++ {

		startindx := int(i-32) * 9
		endindx := startindx + 8

		banner[i] = fileLines[startindx:endindx]
	}

	return banner, nil
}

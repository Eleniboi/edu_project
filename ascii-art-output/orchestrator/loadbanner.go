package orchestrator


import (
	"fmt"
	"os"
	"strings"
)

func Loadbanner(banner string) (map[rune][]string, error) {

	var bannerFont = make(map[rune][]string)

	file, err := os.ReadFile(banner)

	if err != nil {

		return nil, fmt.Errorf("Error reading file: %w", err)
	}

	fileLines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	if len(fileLines) != 856 {
		return nil, fmt.Errorf("invalid file content")
	}

	fileLines = fileLines[1:]

	for i := ' '; i < '~'; i++ {

		startindx := int(i-32) * 9
		endindx := startindx + 8

		bannerFont[i] = fileLines[startindx:endindx]
	}

	return bannerFont, nil
}

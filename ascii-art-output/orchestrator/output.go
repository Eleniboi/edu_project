package orchestrator

import (
	"fmt"
	"os"
	"strings"
)

func Output() (string, string, string, error) {

	switch {

	case len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--output="):

		Outputfile := strings.ToLower(strings.TrimPrefix(os.Args[1], "--output="))

		inputStr := (os.Args[2])

		bannerfile := strings.ToLower(os.Args[3])

		if bannerfile == "standard" || bannerfile == "thinkertoy" || bannerfile == "shadow" {
			bannerfile += ".txt"
		}

		return Outputfile, inputStr, bannerfile, nil
	}

	return "", "", "", fmt.Errorf(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)

}

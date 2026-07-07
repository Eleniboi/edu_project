package main

import (
	"fmt"
	"os"
	"strings"
)

func parser() (string, string, string, error) {

	var fullstr string
	var substr string
	var color string

	switch {
	case len(os.Args) == 2:

		fullstr = os.Args[1]
		return "", "", fullstr, nil

	case len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--color="):

		color = strings.TrimPrefix(os.Args[1], "--color=")
		fullstr = os.Args[2]
		return color, "", fullstr, nil

	case len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--color="):

		color = strings.TrimPrefix(os.Args[1], "--color=")
		substr = os.Args[2]
		fullstr = os.Args[3]
		return color, substr, fullstr, nil

	}

	return color, substr, fullstr, fmt.Errorf("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")

}

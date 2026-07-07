package main

import (
	"strings"
)

func split(input string) []string {

	result := strings.Split(input, "\\n")
	return result
}

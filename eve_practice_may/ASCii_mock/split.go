package main

import "strings"

func SplitInput(input string) []string {

	result := strings.Split(input, "\\n")
	return result
}

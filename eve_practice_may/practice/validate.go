package main

import "fmt"

func validate(input string) (rune, error) {

	var result rune
	for _, r := range input {

		if r < ' ' || r > '~' {
			return 0, fmt.Errorf(", %c is not a printable ascii character", r)
		}
		result = r
	}
	return result, nil
}
package main

import "fmt"

func ValidateInput(input string) (rune, error) {

	for _, ch := range input {

		if ch < ' ' || ch > '~' {
			return ch, fmt.Errorf("%c is not a printable character", ch)
		}

	}
	return 0, nil
}

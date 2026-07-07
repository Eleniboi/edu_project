package main

import (
	"errors"
	"strings"
)

func ValidatInput(input string) error {

	input = strings.ReplaceAll(input, "\r\n", "\n")

	splitInput := strings.Split(input, "\n")

	for _, word := range splitInput {

		for _, char := range word {

			if char < 32 || char > 126 {

				return errors.New("Invalid Character")
			}
		}
	}
	return nil
}

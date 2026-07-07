package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Start from 1 to skip the program name in os.Args[0]
	args := os.Args[1:]

	// Loop through the arguments in reverse order
	for i := len(args) - 1; i >= 0; i-- {
		// Print each argument, character by character
		for _, rune := range args[i] {
			z01.PrintRune(rune)
		}
		// Print a new line after each argument
		z01.PrintRune('\n')
	}
}

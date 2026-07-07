package main

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateBanner(banner map[rune][]string) error {

	if len(banner) != 95 {
		return errors.New("Incomplete printable character")
	}

	for key, value := range banner {

		if key < ' ' || key > '~' {

			return fmt.Errorf("%c is not a printable ascii character", key)
		}

		if len(value) != 8 {
			return errors.New("Ascii Character must be 8 rows")
		}
	}
	return nil
}

func MergeBanners(base, priority map[rune][]string) map[rune][]string {

	banner := map[rune][]string{}

	for key, value := range base {

		banner[key] = value
	}

	for key, value := range priority {

		banner[key] = value
	}
	return banner
}

func TrimArtRows(rows []string) []string {

	var result = make([]string, len(rows))

	for i, ch := range rows {

		result[i] = strings.TrimRight(ch, " ")
	}
	return result
}

func PadArtRows(rows []string, width int) []string {

	var result = make([]string, len(rows))

	for i, ch := range rows {

		current := len(ch)

		if width <= 0 || current >= width {
			result[i] = ch

		} else if len(ch) < width {

			ch = ch + strings.Repeat(" ", width-current)

			result[i] = ch
		}
	}
	return result
}

func StackTwo(top []string, bottom []string) []string {

	result := make([]string, len(top)+len(bottom))

	copy(result, top)
	copy(result[len(top):], bottom)
	return result
}

func StackAll(block [][]string)[]string{

	result := []string{}

	for _, ch := range block{

		result = append(result, ch...)
	}
	return result
}

func main() {

	rows := []string{"sam", "lo"}

	fmt.Println(PadArtRows(rows, -1))
}

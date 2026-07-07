package orchestrator

import "fmt"

var COLORS = map[string]string{

	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

const RESET = "\033[0m"

func Getcolor(colorName string) (string, error) {

	value, ok := COLORS[colorName]

	if !ok {
		return "", fmt.Errorf("%q this color type does not exist", colorName)
	}
	return value, nil
}

package main

import (
	"fmt"
	"os"
	"strings"

	"texttool/transformer"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage: go run . sample.txt result.txt")
		return
	}

	inputFile := args[1]
	outputFile := args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file")
		return
	}

	text := string(data)

	tokens := strings.Fields(text)

	tokens = transformer.Transformer(tokens)

	tokens = transformer.FixArticles(tokens)

	tokens = transformer.FixQuotes(tokens)

	tokens = transformer.FixPunctuation(tokens)

	result := strings.Join(tokens, " ")

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file")
		return
	}
}

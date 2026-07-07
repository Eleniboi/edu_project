package main

import (
	"bufio"
	"fmt"
	"go-reloaded/texttool"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Invalid input: usage go run . input.txt output.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	INfile, err := os.Open(inputfile)

	if err != nil {
		fmt.Println("Error opening file,", err)
		return
	}

	defer INfile.Close()

	OUTfile, err := os.Create(outputfile)

	if err != nil {
		fmt.Println("Error creating file, ", err)
		return
	}

	defer OUTfile.Close()

	scanner := bufio.NewScanner(INfile)
	writer := bufio.NewWriter(OUTfile)

	for scanner.Scan() {

		line := scanner.Text()
		line = texttool.Command(line)
		line = texttool.CommandN(line)
		line = texttool.ToDecimal(line)
		line = texttool.Article(line)
		line = texttool.Punctuation(line)
		line = texttool.Quote(line)

		_, err := writer.WriteString(line + "\n")

		if err != nil {
			fmt.Println("Error writing file, ", err)
			return
		}
	}
	writer.Flush()
}

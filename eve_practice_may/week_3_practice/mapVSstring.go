package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var text string

	fmt.Print("Enter sentence: ")

	scanner.Scan()

	text = scanner.Text()

	splitInput := strings.Split(text, " ")

	var m = make(map[string]int)

	for _, ch := range splitInput {

		for i := range ch {

			if strings.ContainsAny(ch, ":.,?!;") && i == len(ch)-1 {

				ch = ch[:len(ch)-1]
			}
		}
		m[strings.ToLower(ch)]++

	}

	fmt.Println(m)

}

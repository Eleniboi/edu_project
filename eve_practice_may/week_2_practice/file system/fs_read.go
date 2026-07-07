package main

import (
	"bufio"
	"fmt"
	"os"
)

func readWithOs() {

	file, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Printf("error reading file: %w", err)
		return
	}

	fmt.Print(string(file))
}

func readWithBufio() {

	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("error opening file: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var count = 1
	for scanner.Scan() {

		fmt.Println(count, ": ", scanner.Text())
		count++
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("error scanning file: ", err)
		return
	}
}

func main() {
	//readWithOs()
	readWithBufio()
}

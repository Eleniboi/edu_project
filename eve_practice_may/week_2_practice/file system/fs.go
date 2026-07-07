package main

import (
	"bufio"
	"fmt"
	"os"
)

func fs1() {

	// this function ask user for name, write that name gotten to a file and still greet the user with the name,
	// there are two ways of collecting the name:
	// 1. using Scanln: but it is limited as it tend to stop at the first white space
	// 2. using bufio: this allow you to read data in chunck

	var name string

	fmt.Print("Enter your name: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	name = scanner.Text()

	err := os.WriteFile("user.txt", []byte(name), 0644)

	if err != nil {
		fmt.Println("error writing file, ", err)
		return
	}

	data, err := os.ReadFile("user.txt")

	if err != nil {
		fmt.Println("error reading file, ", err)
		return
	}

	fmt.Printf("Hello, %s", string(data))

}

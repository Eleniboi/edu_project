package main

import (
	"fmt"
	"os"
	
)

func main() {

	input := os.Args[1]

	bannerFont := os.Args[2]

	banner, err := LoadBanner(bannerFont)

	if err != nil {
		fmt.Println("Error loading banner")
		return
	}

	result := GenerateArt(input, banner)

	fmt.Print(result)
}

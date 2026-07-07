package main

import (
	orc "asciicolor/orchestrator"
	"fmt"
)

func main() {

	color, substr, fullstr, err := parser()

	if err != nil {
		fmt.Println(err)
		return
	}

	var colorCode string

	if color != "" {
		colorCode, err = orc.Getcolor(color)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// result := colorText(substr, fullstr, colorCode)
	// fmt.Println(result)

	banner, _ := orc.LoadBanner("banner/standard.txt")

	final := orc.GenerateArt(fullstr, substr, colorCode, banner)

	fmt.Print(final)

}

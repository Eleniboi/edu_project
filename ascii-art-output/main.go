package main

import (
	"fmt"
	"os"
	orc "pract/orchestrator"
)

func main() {

	outputFile, inputStr, bannerFile, err := orc.Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	bannerContent, err := orc.Loadbanner("bannerFiles/" + bannerFile)

	if err != nil {
		fmt.Print(err)
	}

	result := orc.GenerateArt(inputStr, bannerContent)

	err = os.WriteFile(outputFile, []byte(result), 0644)
}

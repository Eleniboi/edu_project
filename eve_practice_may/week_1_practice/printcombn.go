package main

import (
	"fmt"
)

func main() {

	var score int
	var currentscore []int
	var total int

	for {
		fmt.Print("input your score: ")
		fmt.Scanln(&score)

		if score < 0 {
			break
		} else {
			currentscore = append(currentscore, score)
			total += score
		}
	}

	var average float64
	if len(currentscore) > 0{

		count := len(currentscore)

		average = float64(total) / float64(count)
	}

	fmt.Println("currentscore: ", currentscore)
	fmt.Println("total: ", total)
	fmt.Printf("average: %2.1f", average)
}

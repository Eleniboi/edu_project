package main

import (
	"fmt"
)

func printComb(){

	for i := 0; i <= 9 ; i++{
		for j := i+1; j <= 9; j++{
			for p := j+1; p <= 9; p++{
				fmt.Printf("%d%d%d, ",i,j,p)
			}
		}
	}
}
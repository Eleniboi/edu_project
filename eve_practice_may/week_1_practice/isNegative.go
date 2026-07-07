package main

import (
	"fmt"
)

//this function print T if the number passed is a nagative number other wise it print F
func IsNegative(nb int) {

	if nb < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

// func main() {

// 	IsNegative(-9)
// 	IsNegative(1)
// 	IsNegative(0)
// 	IsNegative(-1)
// }

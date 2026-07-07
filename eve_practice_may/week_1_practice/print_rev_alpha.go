package main 

import (
	"fmt"
)

func print_rev() {

	for i := 'z'; i >= 'a'; i--{
		fmt.Printf("%c", i)
	}
}

// func main() {
// 	print()
// }
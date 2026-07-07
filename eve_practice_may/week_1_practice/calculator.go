package main

import (
	"fmt"
)

func Help() {

	fmt.Println("For Smooth Experience Follow This Guild Line")
}

/*func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\033[32;1m        WELCOME TO THE CLI CALCULATOR\033[0m")
	time.Sleep(3* time.Second)
	fmt.Println("\033[1mEnter help for Guild line or Enter quit to exit\033[0m")


	for {
		fmt.Print("\033[1m> \033[0m")

		scanner.Scan()

		input := scanner.Text()

		if input == "help"{
			Help()
		}
	}
}
*/

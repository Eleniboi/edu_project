package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func help() {
	fmt.Println("For Smooth Experience enter the following")
	time.Sleep(2 * time.Second)
	fmt.Print("\rconvert  1E hex  ---> 30")
	time.Sleep(2 * time.Second)
	fmt.Print("\rconvert  1010 bin ---> 10")
	time.Sleep(2 * time.Second)
	fmt.Print("\rEnter help for this same menu")
	time.Sleep(2 * time.Second)
	fmt.Print("\rEnter quit to exit the program")
	time.Sleep(2 * time.Second)
	fmt.Print("\rEnter quit to exit the program.")
	time.Sleep(300 * time.Millisecond)
	fmt.Print("\rEnter quit to exit the program..")
	time.Sleep(300 * time.Millisecond)
	fmt.Print("\rEnter quit to exit the program...")

}

func base() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("welcome to base converter")
	fmt.Print(`Enter help for Guildline or quit to exit the program: `)
	//fmt.Print(": ")
	for scanner.Scan() {

		fmt.Print(" > ")

		input := scanner.Text()

		if input == "" {
			fmt.Println("Enter help for Guildline")
			continue
		}
		if input == "help" {
			help()
			continue
		}
		if input == "quit" {
			fmt.Println("GoodBye!!")
			break
		}

		part := strings.Fields(input)

		if len(part) != 3 {

			fmt.Println("invalid input: type help for guideline")
			continue
		}

		convert := part[0]

		if convert != "convert" {
			fmt.Println("invalid input Enter help for guildline")
			continue
		}

		BaseNum := part[1]
		Base := part[2]

		switch strings.ToLower(Base) {
		case "hex":
			n, err := strconv.ParseInt(BaseNum, 16, 64)
			if err != nil {
				fmt.Printf("%s is not a valid hex number", BaseNum)
				continue
			}
			fmt.Println("Decimal = ", n)
		case "bin":
			n, err := strconv.ParseInt(BaseNum, 2, 64)

			if err != nil {
				fmt.Printf("%s is not a valid bin number\n", err)
				continue
			}
			fmt.Println("Decimal: ", n)
		}
	}
}

func main() {
	base()
}

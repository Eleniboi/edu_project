package main

import (
	"fmt"
	"strings"
)

func Checker(name string) string {

	var AvailableN = map[string]bool{

		"samuel": true,
		"john":   true,
		"omafu":  true,
		"james":  true,
	}

	newName := strings.ToLower(name)

	_, ok := AvailableN[newName]

	if ok {
		fmt.Sprintln("Username taken!! \nAvailable username: \n"+name+"1", "\n"+name+"2", "\n"+name+"3")

	}

	AvailableN[newName] = true
	return fmt.Sprintln("user name is successfully updated")

}

func main() {

	var userName string

	fmt.Print("Enter UserName: ")

	fmt.Scanln(&userName)

	fmt.Print(Checker(userName))
}

//and if it does not exist i will add it to the map like this newName = AvailableN[newName]

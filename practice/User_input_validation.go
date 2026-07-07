package main

import "fmt"

func main() {

	var name string
	var birthday int
	var currentyear = 2026

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Birthyear: ")
	fmt.Scanln(&birthday)

	age := currentyear - birthday

	if age >= 50 {
		fmt.Println(name, "you are", age, "years old", "and the required age range from 18 - 40")
		return
	} else if age >= 18 {
		fmt.Println("Hello", name, "you are", age, "years Old")
	} else if age <=17{
		fmt.Println(name,"you are",age,"valid registration age 18 Above")
		return 
	}

	var school string
	var church string
	var work string

	fmt.Print("Enter School Name: ")
	fmt.Scanln(&school)

	fmt.Print("Enter Church Name: ")
	fmt.Scanln(&church)

	fmt.Print("Your Occupation: ")
	fmt.Scanln(&work)

	fmt.Println(name, "you are", age, "years old working as a/an", work, "you attended", school, "and worship at", church, "church you have successfully completed your registration and", age, "years is an accepted age for this course")
}

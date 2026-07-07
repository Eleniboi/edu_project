package main

import (
	"bufio"
	"fmt"
	"os"
)

// Create file
func Createfile() {

	// this function create a file Myname.txt and write a text to that same file
	newfile, err := os.Create("Myname.txt")

	if err != nil {
		fmt.Println("error creating file")
		return
	}
	defer newfile.Close()

	fmt.Print("Write your full name: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	myname := scanner.Text()

	err = os.WriteFile("Myname.txt", []byte(myname), 0644)

	if err != nil {
		fmt.Println("error writing file", err)
		return
	}
}

// this function rename file
func renamefile() {

	file, err := os.Create("user.txt")

	if err != nil {
		fmt.Println("error creating file")

		return
	}
	defer file.Close()

	err = os.Rename("user.txt", "final.txt")

	if err != nil{
		fmt.Println("error renaming file, ", err)
		return 
	}
	fmt.Println("file Renamed successfully")
}

func deleteFile(name string){


	err := os.Remove(name)

	if err != nil{
		fmt.Println("file not found")
		return 
	}
	fmt.Println("file deleted successfully")
}

func main() {

	Createfile()
	renamefile()
}
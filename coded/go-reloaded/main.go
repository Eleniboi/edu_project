package main 


import(
	"fmt"
	"os"
)


func main() {

	if len(os.Args) != 3{
		fmt.Println("invalid input: usage go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	file, err := os.ReadFile(inputFile)

	if err != nil{
		fmt.Println(err)
		return
	}

	text := string(file)
	text = quote(text)
	text = ToDecimal(text)
	text = punct(text)

	err = os.WriteFile(outputFile, []byte(text), 0644)

	if err != nil{
		fmt.Println("error writing file")
		return
	}
}
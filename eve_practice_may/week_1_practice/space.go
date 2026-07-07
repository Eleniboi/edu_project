package main

func space(s string) string {

	var result string
	for _, r := range s {

		result += string(r) + " "
	}
	return result
}

// func main() {

// 	fmt.Println(space("samuel"))
// }

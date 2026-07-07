package main

import (
	"fmt"
)

func printComb2() {
	for i := 0; i <= 98; i++ {
		for j := 1; j <= 99; j++ {
			fmt.Printf("%02d %02d, ", i, j)
			
		}
	}
}

// func main(){

// 	printComb2()
// }

//////////////////////////////////////////////////////////////////
// func printComb2(s string) string {
// 	text := strings.Fields(s)

// 	for i := 0; i < len(text); i++ {

// 		if strings.HasPrefix(text[i], "(") || strings.HasSuffix(text[i], ")") && strings.HasSuffix(text[i+1], ")") {

// 			cmd := strings.TrimPrefix(text[i], "(")
// 			cmd = strings.TrimSuffix(cmd, ")")
// 			cmd = strings.TrimSuffix(cmd, ",")

// 			clean := strings.TrimSuffix(text[i+1], ")")

// 			text = append(text[:i], text[i+2:]...)
// 			i--

// 			num, _ := strconv.Atoi(clean)

// 			for j := i - num + 1; j <= i; j++ {
// 				if cmd == "up" {
// 					text[j] = strings.ToUpper(text[j])
// 				}
// 				if cmd == "low" {
// 					text[j] = strings.ToLower(text[j])
// 				}
// 				if cmd == "cap" {
// 					text[j] = strings.ToUpper(text[j][:1]) + strings.ToLower(text[j][1:])
// 				}
// 				if cmd == "hex"{
// 					n, _ := strconv.ParseInt(text[j], 16 , 64)
// 					text[j] = fmt.Sprint(n)
// 				}
// 			}
// 		}
// 	}
// 	return strings.Join(text, " ")
// }

// func main() {

// 	fmt.Println(printComb2("samuel 1E (hex) you are great okay (up, 3)"))
// }

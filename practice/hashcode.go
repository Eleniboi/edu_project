package main

import "fmt"

func HashCode(dec string) string {
	res := ""
	for _, r := range dec {
		h := (int(r) + len(dec)) % 127
		if h < 33 {
			res += string(rune(h + 33))
		} else {
			res += string(rune(h))
		}
	}
	return res
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

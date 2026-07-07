package main

func firstWord(s string) string {

	for i, ch := range s {

		if ch == ' ' {
			return s[:i]
		}
	}
	return s
}
func First_word(s string) string {

	//p := " "
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			return s[:i]
		}
	}
	return s
}

// func main() {

// 	fmt.Print(First_word("samuel you are great oo"))
// }

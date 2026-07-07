package main



// this function return the recieves Ascii input and return the ascii character and it ascii value on the table
func AsciiValue(s string) map[string]int {

	var maper = make(map[string]int)

	for i, r := range s {
		maper[string(s[i])] = int(r)
	}
	return maper
}

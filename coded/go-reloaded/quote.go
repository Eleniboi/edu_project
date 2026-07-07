package main


import(
	"strings"
)


func quote(s string) string {

	text := strings.Split(s, "'")

	for i := 0; i < len(text); i++{

		if i % 2 == 1{
			text[i] = strings.TrimSpace(text[i])
		}
	}
	return strings.Join(text, "'")
}

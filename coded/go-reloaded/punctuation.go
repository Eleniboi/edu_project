package main


import(
	"regexp"
)


func punct(s string) string {


	pat := regexp.MustCompile(`\s+([:;?!.,]+)`)

	s = pat.ReplaceAllString(s, "$1")

	pat2 := regexp.MustCompile(`([:;?!.,]+)([^:,.;?!])`)

	s = pat2.ReplaceAllString(s, "$1 $2")
	

	return s
}
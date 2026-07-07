package main

import (
	"strings"
)

func ReverseRows(rows []string) []string {

	var result = make([]string, len(rows))

	copy(result, rows)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {

		result[i], result[j] = result[j], result[i]
	}
	return result
}
func Double(numbers []int) []int {

	for i := range numbers {
		numbers[i] *= 2
	}

	return numbers
}

func tran(s string) []string {

	result := []string{}

	for _, ch := range s {

		result = append(result, string(ch))
	}
	return result
}

func CleanWords(words []string) []string {

	result := []string{}

	for _, word := range words {

		clean := strings.ToLower(strings.TrimSpace(word))

		if clean != "" {
			result = append(result, clean)
		}
	}
	return result
}

// func main() {

// 	fmt.Println(tran("samuel"))
// 	nums := []int{1, 2, 3}

// 	newNums := Double(nums)

// 	fmt.Println(newNums)
// 	fmt.Println(CleanWords([]string{
// 		"",
// 		"Lang",
// 		"Fun",
// 		"go",
// 		"     ",
// 	}))
// }

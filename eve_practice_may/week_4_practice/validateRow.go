package main

import "fmt"

func ValidateRows(rows []string, maxWidth int) error {

	for i, ch := range rows {

		switch {
		case ch == "":
			return fmt.Errorf("row %d is empty", i+1)
		case len(ch) > maxWidth:
			return fmt.Errorf("len of rows must not be greater than the maxWidth %d", maxWidth)
		}
	}
	return nil
}

func MergeRows(left, right []string) []string {

	var bothsides []string

	for i := range right {

		bothsides = append(bothsides, left[i]+right[i])
	}
	return bothsides
}

func FlattenWords(words [][]string) []string {

	var result []string

	for _, ch := range words {

		result = append(result, ch...)
	}
	return result
}

func GroupByLength(words []string) map[int][]string{

	var result = make(map[int][]string)


	for _, ch := range words{
		length := len(ch)
			result[length] = append(result[length], ch)
	}
	return result
}

func main() {

	// left := []string{"AAA", "BBB", "CCC"}
	// right := []string{"111", "222", "333"}

	// result := MergeRows(left, right)

	// for i := range result {

	// 	fmt.Print(result[i], "\n")
	// }
	// fmt.Print(len(result))

	// slice := [][]string{{"AAA", "BBB"}, { "CCC"},}

	// fmt.Println(FlattenWords(slice))
	// fmt.Printf("%#v", slice)

	fmt.Println(GroupByLength([]string{"how","he","bro","to"}))

}

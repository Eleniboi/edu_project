package transformer

import (
	"strconv"
	"strings"
)

func Transformer(tokens []string) []string {

	var result []string

	for i := 0; i < len(tokens); i++ {

		if tokens[i] == "(bin)" {
			convertPrevious(result, 2)

		} else if tokens[i] == "(hex)" {
			convertPrevious(result, 16)

		} else if tokens[i] == "(up)" {
			upperPrevious(result, 1)

			          
		} else if tokens[i] == "(up," && i < len(tokens)-1 {

			clean := strings.TrimSuffix(tokens[i+1], ")")
			n, _ := strconv.Atoi(clean)

			upperPrevious(result, n)

			i++

			
		} else if strings.HasPrefix(tokens[i], "(up,") {

			clean := strings.TrimPrefix(tokens[i], "(up,")
			clean = strings.TrimSuffix(clean, ")")
			n, _ := strconv.Atoi(clean)

			upperPrevious(result, n)

		} else if tokens[i] == "(low)" {
			lowerPrevious(result, 1)

			
		} else if tokens[i] == "(low," && i < len(tokens)-1 {

			clean := strings.TrimSuffix(tokens[i+1], ")")
			n, _ := strconv.Atoi(clean)

			lowerPrevious(result, n)

			i++

			
		} else if strings.HasPrefix(tokens[i], "(low,") {

			clean := strings.TrimPrefix(tokens[i], "(low,")
			clean = strings.TrimSuffix(clean, ")")
			n, _ := strconv.Atoi(clean)

			lowerPrevious(result, n)

		} else if tokens[i] == "(cap)" {
			capPrevious(result, 1)

			
		} else if tokens[i] == "(cap," && i < len(tokens)-1 {

			clean := strings.TrimSuffix(tokens[i+1], ")")
			n, _ := strconv.Atoi(clean)

			capPrevious(result, n)

			i++

		} else if strings.HasPrefix(tokens[i], "(cap,") {

			clean := strings.TrimPrefix(tokens[i], "(cap,")
			clean = strings.TrimSuffix(clean, ")")
			n, _ := strconv.Atoi(clean)

			capPrevious(result, n)

		} else {
			result = append(result, tokens[i])
		}
	}

	return result
}

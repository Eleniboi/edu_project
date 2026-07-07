package texttool

import (
	"fmt"
	"strconv"
	"strings"
)

func ToDecimal(s string) string {

	result := strings.Fields(s)

	for i := 0; i < len(result); i++ {

		switch strings.ToLower(result[i]) {
		case "(hex)":
			if i > 0 {
				n, _ := strconv.ParseInt(result[i-1], 16, 64)
				result[i-1] = fmt.Sprint(n)
				result = append(result[:i], result[i+1:]...)
				i--
			}
		case "(bin)":
			if i > 0 {
				n, _ := strconv.ParseInt(result[i-1], 2, 64)
				result[i-1] = fmt.Sprint(n)
				result = append(result[:i], result[i+1:]...)
				i--
			}
		}

	}
	return strings.Join(result, " ")
}

// if strings.HasPrefix(strings.ToLower(result[i]), "(hex"){
// // 	n, _ := strconv.ParseInt(result[i-1], 16, 64)

// // 	result[i-1] = fmt.Sprint(n)
// // 	result = append(result[:i], result[i+1:]...)
// // 	i--

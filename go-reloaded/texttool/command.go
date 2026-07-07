package texttool

import (
	"strings"
)

func Command(s string) string {

	result := strings.Fields(s)

	for i := 0; i < len(result); i++ {

		switch strings.ToLower(result[i]) {
		case "(up)":
			if i > 0 {
				result[i-1] = strings.ToUpper(result[i-1])
				result = append(result[:i], result[i+1:]...)
				i--
			}
		case "(low)":
			if i > 0 {
				result[i-1] = strings.ToLower(result[i-1])
				result = append(result[:i], result[i+1:]...)
				i--
			}
		case "(cap)":
			if i > 0 {
				result[i-1] = strings.ToUpper(string(result[i-1][:1])) + strings.ToLower(result[i-1][1:])
				result = append(result[:i], result[i+1:]...)
				i--
			}
		}

	}
	return strings.Join(result, " ")
}

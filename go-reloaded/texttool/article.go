package texttool

import (
	"strings"
)

func Article(s string) string {

	result := strings.Fields(s)
	for i := 0; i < len(result); i++ {

		if i+1 < len(result) {
			switch result[i] {
			case "a":
				if strings.ContainsAny("aeiouhAEIOUH", string(result[i+1][:1])) {
					result[i] = "an"
				}
			case "A":
				if strings.ContainsAny("aeiouhAEIOUH", string(result[i+1][:1])) {
					result[i] = "An"
				}
			case "an":
				if !strings.ContainsAny("aeiouhAEIOUH", string(result[i+1][:1])) {
					result[i] = "a"
				}
			case "An":
				if !strings.ContainsAny("aeiouhAEIOUH", string(result[i+1][:1])) {
					result[i] = "A"
				}
			}
		}
	}
	return strings.Join(result, " ")
}

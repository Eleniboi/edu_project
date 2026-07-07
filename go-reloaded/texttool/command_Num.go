package texttool

import (
	"strconv"
	"strings"
)

func CommandN(s string) string {

	result := strings.Fields(s)

	for i := 1; i < len(result); i++ {
		if result[i] == "(up," && i+1 < len(result) {
			clean := strings.TrimSuffix(result[i+1], ")")

			num, _ := strconv.Atoi(clean)
			result = append(result[:i], result[i+2:]...)
			i--

			start := i - num + 1

			if start < 0 {
				start = 0
			}

			for j := start; j <= i; j++ {
				result[j] = strings.ToUpper(result[j])
			}
		}
		if strings.HasPrefix(result[i], "(up,") && strings.HasSuffix(result[i], ")") {
			clean := strings.TrimSuffix(strings.TrimPrefix(result[i], "(up,"), ")")

			num, _ := strconv.Atoi(clean)
			result = append(result[:i], result[i+1:]...)
			i--

			start := i - num + 1
			if start < 0 {
				start = 0
			}

			for j := start; j <= i; j++ {
				result[j] = strings.ToUpper(result[j])
			}
		}
		if result[i] == "(low," && i+1 < len(result) {
			clean := strings.TrimSuffix(result[i+1], ")")

			num, _ := strconv.Atoi(clean)
			result = append(result[:i], result[i+2:]...)
			i--

			start := i - num + 1

			if start < 0 {
				start = 0
			}

			for j := start; j <= i; j++ {
				result[j] = strings.ToLower(result[j])
			}
		}
		if strings.HasPrefix(result[i], "(low,") && strings.HasSuffix(result[i], ")") {
			clean := strings.TrimPrefix(strings.TrimSuffix(result[i], ")"), "(low,")

			num, _ := strconv.Atoi(clean)
			result = append(result[:i], result[i+1:]...)

			count := i - num + 1

			if count < 0 {
				count = 0
			}
			for j := count; j <= i; j++ {
				result[j] = strings.ToLower(result[j])
			}
		}
		if result[i] == "(cap," && i+1 < len(result) {
			clean := strings.TrimSuffix(result[i+1], ")")

			num, _ := strconv.Atoi(clean)
			result = append(result[:i], result[i+2:]...)
			i--

			count := i - num + 1

			if count < 0 {
				count = 0
			}

			for j := count; j <= i; j++ {
				result[j] = strings.ToUpper(result[j][:1]) + strings.ToLower(result[j][1:])
			}
		}
		if strings.HasPrefix(result[i], "(cap,") && strings.HasSuffix(result[i], ")") {
			clean := strings.TrimPrefix(result[i], "(cap,")
			clean = strings.TrimSuffix(clean, ")")

			num, _ := strconv.Atoi(clean)
			result = append(result[:i], result[i+1:]...)
			i--

			count := i - num + 1

			if count < 0 {
				count = 0
			}

			for j := count; j <= i; j++ {
				result[j] = strings.ToUpper(result[j][:1]) + strings.ToLower(result[j][1:])
			}
		}

	}
	return strings.Join(result, " ")
}

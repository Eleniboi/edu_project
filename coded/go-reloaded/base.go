package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CapFirst(s string) string {

	//but it is better to use strings.Builder more efficient it has an inner bufer
	var result string
	for i, ch := range s {

		if ch >= 'a' && ch <= 'z' && i == 0 {

			result += string(ch - 32)
			continue
		}
		result += string(ch)
	}

	return result
}

func ToDecimal(s string) string {

	text := strings.Fields(s)

	for i := 0; i < len(text); i++ {

		if strings.ToLower(text[i]) == "(hex)" && i > 0 {

			num, err := strconv.ParseInt(text[i-1], 16, 64)

			if err != nil {
				return fmt.Sprint(err)
			}
			text[i-1] = fmt.Sprint(num)
			text = append(text[:i], text[i+1:]...)
			i--
		}

		if strings.ToLower(text[i]) == "(low)" && i > 0 {

			text[i-1] = strings.ToLower(text[i-1])
			text = append(text[:i], text[i+1:]...)
			i--
		}

		if strings.ToLower(text[i]) == "(cap)" {

			text[i-1] = CapFirst(text[i-1])
			text = append(text[:i], text[i+1:]...)
			i--
		}
		if strings.HasPrefix(text[i], "(up,") && strings.HasSuffix(text[i], ")") {

			clean := strings.TrimPrefix(text[i], "(up,")
			clean = strings.TrimSuffix(clean, ")")

			num, _ := strconv.Atoi(clean)

			text = append(text[:i], text[i+1:]...)
			i--

			count := max(i-num+1, 0)

			for j := count; j <= i; j++ {
				text[j] = strings.ToUpper(text[j])
			}
		}
	}
	return strings.Join(text, " ")
}

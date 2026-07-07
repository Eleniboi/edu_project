package transformer

import "strconv"

func convertPrevious(result []string, base int) {
	if len(result) == 0 {
		return
	}

	lastIndex := len(result) - 1
	lastWord := result[lastIndex]

	value, err := strconv.ParseInt(lastWord, base, 64)
	if err != nil {
		return
	}
	result[lastIndex] = strconv.FormatInt(value, 10)
}

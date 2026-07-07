package transformer

func FixQuotes(tokens []string) []string {

	var result []string
	insideQuote := false

	for i := 0; i < len(tokens); i++ {

		if tokens[i] == "'" {

			if !insideQuote && i < len(tokens)-1 {
				tokens[i+1] = "'" + tokens[i+1]
			}

			if insideQuote && len(result) > 0 {
				result[len(result)-1] = result[len(result)-1] + "'"
			}

			insideQuote = !insideQuote
			continue
		}

		result = append(result, tokens[i])
	}

	return result
}

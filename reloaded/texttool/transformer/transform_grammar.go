package transformer

func FixArticles(tokens []string) []string {

	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'h': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'H': true,
	}
	for i := 0; i < len(tokens); i++ {

		if (tokens[i] == "a" || tokens[i] == "A") && i < len(tokens)-1 {

			first := tokens[i+1][0]

			if vowels[first] {

				if tokens[i] == "A" {
					tokens[i] = "An"
				} else {
					tokens[i] = "an"
				}
			}
		}
	}
	return tokens
}
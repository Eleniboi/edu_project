package main

func Rot14(s string) string {
	result := []rune(s)
	for i, r := range result {

		if r >= 'A' && r <= 'Z' {
			result[i] = (result[i] - 'A' + 14) % 26
			result[i] = result[i] + 'A'
		}
	}
	return string(result)
}

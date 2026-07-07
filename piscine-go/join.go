package piscine

func Join(strs []string, sep string) string {
	sliceString := []string(strs)
	result := strs[0]
	for i := 1; i < len(sliceString); i++ {
		result = result + sep + sliceString[i]
	}
	return result
}

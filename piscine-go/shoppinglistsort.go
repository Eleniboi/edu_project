package piscine

func ShoppingListSort(slice []string) []string {
	s := len(slice)

	for i := 0; i < s-1; i++ {
		for j := 0; j < s-i-1; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}

	return slice
}

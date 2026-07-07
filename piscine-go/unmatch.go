package piscine

func Unmatch(a []int) int {
	count := make(map[int]int)

	for _, i := range a {
		count[i]++
	}

	for _, i := range a {
		if count[i]%2 != 0 {
			return i
		}
	}
	return -1
}

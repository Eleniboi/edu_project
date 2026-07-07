package main


//this function check if a particular length of row is less than the width then it counts
func CountShortRows(rows []string, width int) int {

	count := 0

	for _, ch := range rows {

		if len(ch) < width && ch != "" {
			count++
		}
	}
	return count
}

// func main() {

// 	rows := []string{"howfar", "", "hi", "good"}
// 	fmt.Println(CountShortRows(rows, 6))
// }
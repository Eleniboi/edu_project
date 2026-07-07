package main



//this function forces you to learn comparison, and how to fine the higest value of something or the lowest value
//it checks for the longest row
func  LongestRow(rows []string) string {

	if len(rows) == 0{
		return ""
	}

	//
	longest := rows[0]

	for _, ch := range rows {

		if len(ch) > len(longest) {
			longest = ch
		}
	}
	return longest
}
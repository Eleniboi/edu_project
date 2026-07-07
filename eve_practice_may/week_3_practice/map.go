package main

import "fmt"

// this function help to build your knowledge in map,  and here a struct is use as a value type
func map2() {

	type student struct {
		Age   int
		score int
	}

	studentScore := []int{88, 98, 38}
	studentName := []string{"samuel", "victor", "james", "gab"}
	studentAge := []int{18, 14, 30, 22, 16, 16}

	var result = make(map[string]student)

	for i, ch := range studentName {

		// this is a condition that checks if 'i' is thesame with len of studentScore before adding the value in the struct
		if i < len(studentScore) {
			result[ch] = student{Age: studentAge[i], score: studentScore[i]}
		}
	}
	fmt.Println(result)
}

package main

import (
	"errors"
	"fmt"
)

func AverageScore(score []float64) (float64, error) {

	var Total float64

	if len(score) == 0 {
		return 0, errors.New("Average score is empty")
	}
	var div int

	for i, value := range score {

		if i < len(score) {
			Total += value
		}
		div = i
	}

	Ave := Total / float64(div+1) 
	return Ave, nil
}

func getGrade(grade float64)string{

	if grade > 100{
		return fmt.Sprintln("Score is out of range")
	}
	if grade >= 70 {
		return fmt.Sprintln("A")
	} else if grade >= 60 {
		return fmt.Sprintln("B")
	}else if grade >= 50 {
		return fmt.Sprintln("C")
	}
	return fmt.Sprintln("F")
}


func studentReport(Student_name string, scores []float64)string{

	report, err := AverageScore(scores)

	if err != nil{
		return fmt.Sprintln(err)
	}

	Final := getGrade(report)

	return fmt.Sprintf("%s | Average %.1f | Grade %s ", Student_name, report, Final)
}

// func main() {

// 	fmt.Println()
// 	fmt.Println(studentReport("pepe", []float64{4,4,33,}))
// }

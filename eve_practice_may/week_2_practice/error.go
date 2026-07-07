package main

import (
	"errors"
	"fmt"
)

var ErrZeroDivision = errors.New("can not be divided by 0")

// this function recieves two numbers of type float64, float64 and error it also handle error like dividing a number by 0
func divide(a, b float64) (float64, error) {

	var result float64

	if b == 0 {
		return 0, ErrZeroDivision
	}

	result = a / b

	return result, nil
}

// this function recieves two input which are of type float64, and return float64 and an error, here the error is been wrapped, adding an error message to an already existing error
func calculate_total(price, quantity float64) (float64, error) {

	res, err := divide(price, quantity)

	if err != nil {
		return 0, fmt.Errorf("Failed calculation: %w", err)
	}
	return res, nil
}

func main() {

	// test for calculate_total
	result, Err := calculate_total(20, 0)

	if Err != nil {
		fmt.Println(Err)//check for error
		fmt.Println(errors.Is(Err, ErrZeroDivision))//check if the error is the same division error in the divide function here it print // true 
	}
	fmt.Println("result: ", result)
	// test for divide function
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Invalid, ", err)
		return
	}

	fmt.Println("result", res)
}

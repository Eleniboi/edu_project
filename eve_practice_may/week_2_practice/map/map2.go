package main

import (
	"fmt"
)

func map2() {

	//this function check 

	availablefruit := []string{"banana", "mango", "apple", "banana", "apple", "mango", "banana", "melon"}

	record := make(map[string]int)

	for _, item := range availablefruit {

		record[item]++
	}
	record["sweet"] = 0
	fmt.Println(record)

	list := "sweet"

	amount, exist := record[list]

	if !exist {

		fmt.Printf("we don't sell %s here\n", list)
	} else if amount == 0 {
		fmt.Printf("we sell %s but current out of stuck\n", list)
	} else {
		fmt.Printf("we have %s only %d available\n", list, amount)
	}

	for ch, i := range record {
		fmt.Printf("%s: %d \n", ch, i)
	}
}

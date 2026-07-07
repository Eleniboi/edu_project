package main

import (
	"fmt"
	"os"
)

func main() {

	banner := "standard.txt"
	
	input := os.Args[1]

	bannerFont, _ := LoadBanner(banner)

	final := generateArt(input, bannerFont)

	fmt.Print(final)
}

// input := os.Args[1]

// banner := os.Args[2]

// file, err := os.ReadFile(banner)

// if err != nil {
// 	fmt.Println("Error reading file")
// 	return
// }

// //splitInput := strings.Split(input, "\\n")

// fileLines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

// //fmt.Println(len(fileLines)/9)
// var bannerFont = make(map[rune][]string)
// for i := ' '; i <= '~'; i++ {
// 	startindx := int(i-32) * 9 + 1
// 	endindx := startindx + 8

// 	bannerFont[i] = fileLines[startindx:endindx]
// }

// var result []string
// var build strings.Builder

// 	for i := 0; i < 8; i++ {
// 		for _, r := range input {

// 			build.WriteString(bannerFont[r][i])
// 		}
// 		result = append(result, build.String())
// 		build.Reset()
// 	}
// 	//fmt.Println(result)
// 	var done strings.Builder
// 	var final string
// 	for i := 0; i < len(result); i++{
// 		done.WriteString(result[i]+"\n")
// 		final = done.String()

// 	}
// 	fmt.Println(final)
// 	err = os.WriteFile("result.txt", []byte(final), 0644)
// }

//	err = os.WriteFile("result.txt", []byte())

// var bannerFont = make(map[rune][]string)

// for j := 0; j < 95; j++ {
// 	startindx := j*9 + 1
// 	endindx := startindx + 8
// 	bannerFont[rune(j+32)] = fileLines[startindx:endindx]

// }

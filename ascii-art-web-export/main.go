package main

import (
	"fmt"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ascii-art", asciiHandler)

	http.HandleFunc("/download", downloadHandler)

	http.HandleFunc("/about", aboutHandler)

	fmt.Println("server is live on port 8000...")

	err := http.ListenAndServe(":8000", nil)

	if err != nil{
		fmt.Println(err)
		return 
	}
}

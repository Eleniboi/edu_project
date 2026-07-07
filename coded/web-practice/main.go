package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ascii-art", asciiHandler)

	http.HandleFunc("/about", aboutHandler)

	fmt.Println("server is live on port 8000...")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

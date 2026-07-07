package main

import (
	"net/http"
	"fmt"
)

func homeHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "Hello")
}

func main() {

	http.HandleFunc("/ascii", homeHandler)
	http.ListenAndServe(":8080", nil)
}

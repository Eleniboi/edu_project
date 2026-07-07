package main

import (
	"html/template"
	"net/http"
)

//this exercise teaches me how to use my template action in my html

type pageData struct {
	Name   []string
	Active string
}

func nameHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("index.html"))

	Data := pageData{
		Name:   []string{"samuel", "emmanuel", "jireh", "maryam", "yamz"},
		Active: "they are all active!!!",
	}

	tmpl.Execute(w, Data)
}

// func main() {

// 	http.HandleFunc("/name", nameHandler)
// 	fmt.Println("server is live on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// }

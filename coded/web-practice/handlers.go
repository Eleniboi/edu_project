package main

import (
	"html/template"
	"net/http"
)

type file struct {
	Title      string
	Result     string
	Banner     string
	ResultArea string
}
//the file was parse globally and used in each handlerFunctions 
var tmpl, _ = template.ParseFiles("template/index.html")

//homeHandler displays and handles what ever is in the homepage of a web
func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	Data := file{

		Title:      "ASCII-ART-WEB",
		Result:     "",
		ResultArea: "Result will be displayed here",
	}

	tmpl.Execute(w, Data)
}

//
func asciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner == "" {
		Data := file{

			Title:      "ASCII-ART-WEB",
			Result:     "Select a banner",
			Banner:     banner,
			ResultArea: "Invalid Input!!",
		}
		tmpl.Execute(w, Data)
		return
	}

	result := Artbuilder(text, "banners/"+banner)

	Data := file{
		
		Title:      "ASCII-ART-WEB",
		Result:     result,
		Banner:     banner,
		ResultArea: "Generated Art",
	}

	tmpl.Execute(w, Data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/about" {

		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("template/about.html"))

	tmpl.Execute(w, nil)
}

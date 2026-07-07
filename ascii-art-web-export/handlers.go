package main

import (
	"html/template"
	"net/http"
	"os"
	"fmt"
)

type file struct {
	Text 	   string
	Title      string
	Result     string
	Banner     string
	ResultArea string
}
//the file was parse globally and used in each handlerFunctions 
var tmpl = template.Must(template.ParseFiles("template/index.html"))

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

	if r.Method != http.MethodPost{
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
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
			Text: text,
		}
		tmpl.Execute(w, Data)
		return
	}

	err := ValidatInput(text); if err != nil{
		http.Error(w, "400 Bad Reques, Invalid character!!!", http.StatusBadRequest)
		return
	}

	result, err := Artbuilder(text, banner)

	if err != nil{
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}

	Data := file{
		
		Title:      "ASCII-ART-WEB",
		Result:     result,
		Banner:     banner,
		ResultArea: "Generated Art",
		Text: text,
	}

	tmpl.Execute(w, Data)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {


	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {

		http.Error(w, "400 Bad Request, Ascii art must be generated before downloading", http.StatusBadRequest)
		return
	}

	if r.Method != http.MethodGet{
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/download"{
		http.Error(w, "404 Not found", http.StatusNotFound)
		return  
	}

	result, err := Artbuilder(text, banner)

	fileName := "Ascii-art.txt"

	err = os.WriteFile(fileName, []byte(result), 0644)

	if err != nil{
		http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}


	data, err := os.ReadFile(fileName)

	if err != nil{
	http.Error(w, "500 Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprint(len(data)))
	w.Header().Set("Content-Disposition", `attachment; filename="` + fileName+`"`)

	w.Write(data)

	os.Remove(fileName)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/about" {

		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("template/about.html"))

	tmpl.Execute(w, nil)
}

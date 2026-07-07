package main


import (
	"net/http"
	"fmt"
)

func homeHandler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "<h1>samuel you are great</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/about"{

		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "<P><em>samuel is a great man of God,<br>And God ilkljioios using him to do great did, 'i will never fail God' samuel's words<em><p>")
}
func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("server is live on port 8080...")
	http.ListenAndServe(":8080", mux)
}
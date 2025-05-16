package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

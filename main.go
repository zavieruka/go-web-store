package main

import (
	"html/template"
	"net/http"
)

type Products struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Products{
		{"Product 1", "Description 1", 10.0, 5},
		{"Product 2", "Description 2", 20.0, 10},
		{"Product 3", "Description 3", 30.0, 15},
		{"Product 4", "Description 4", 40.0, 20},
	}

	templates.ExecuteTemplate(w, "Index", products)
}

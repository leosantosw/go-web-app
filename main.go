package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{
			Name:        "Camisa GO",
			Description: "Camisa estilosa e confortável, perfeita para desenvolvedores.",
			Quantity:    10,
			Price:       100.00,
		},
		{
			Name:        "Caneca GO",
			Description: "Caneca de cerâmica com o logo da linguagem Go.",
			Quantity:    20,
			Price:       25.00,
		},
		{
			Name:        "Livro 'A Linguagem de Programação Go'",
			Description: "Livro oficial para aprender Go.",
			Quantity:    30,
			Price:       80.00,
		},
	}
	templates.ExecuteTemplate(w, "Index", products)
}

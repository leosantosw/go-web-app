package main

import (
	"fmt"
	"net/http"
	"text/template"

	"web-app/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := models.ListProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

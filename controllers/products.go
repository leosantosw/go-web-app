package controllers

import (
	"net/http"
	"text/template"

	"web-app/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

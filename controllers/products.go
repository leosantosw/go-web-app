package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"web-app/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListProducts()

	templates.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		productName := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")
		formattedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Price parameter is invalid")
		}

		formattedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Quantity parameter is invalid")
		}
		models.CreateNewProduct(productName, description, formattedPrice, formattedQuantity)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		path := strings.TrimPrefix(r.URL.Path, "/remove/")
		id := strings.Split(path, "/")[0]
		models.RemoveProduct(id)
	}
	http.Redirect(w, r, "/", http.StatusNoContent)
}

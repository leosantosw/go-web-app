package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		fmt.Println(productName, description, price, quantity)
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
	// http.Redirect(w, r, "/", 301)
}

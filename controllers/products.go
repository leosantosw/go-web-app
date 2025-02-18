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
	if r.Method == "GET" {
		templates.ExecuteTemplate(w, "New", nil)
		return
	} else if r.Method == "POST" {
		HandleCreateProduct(r)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := ExtractURLId(r.URL.Path, "/remove/")
		models.RemoveProduct(id)
	}
	http.Redirect(w, r, "/", http.StatusNoContent)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := ExtractURLId(r.URL.Path, "/edit/")
		product := models.GetProductById(id)
		templates.ExecuteTemplate(w, "Edit", product)
		return
	} else if r.Method == "POST" {
		HandleEditProduct(r)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func HandleCreateProduct(r *http.Request) {
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

func HandleEditProduct(r *http.Request) {
	id := r.FormValue("id")
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
	models.EditProduct(id, productName, description, formattedPrice, formattedQuantity)
}

func ExtractURLId(path string, prefix string) string {
	parts := strings.TrimPrefix(path, prefix)
	return strings.Split(parts, "/")[0]
}

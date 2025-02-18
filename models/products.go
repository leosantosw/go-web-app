package models

import (
	"time"

	"web-app/database"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	CreatedAt   time.Time
}

func ListProducts() []Product {
	db := database.Connection()

	var products []Product

	result, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		product := Product{}
		err = result.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := database.Connection()

	productQuery, err := db.Prepare("INSERT INTO PRODUCTS (name, description, price, quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	productQuery.Exec(name, description, price, quantity)
	defer db.Close()
}

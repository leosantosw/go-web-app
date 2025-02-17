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

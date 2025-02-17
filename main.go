package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "leosantosw"
	password = "leosantosw"
	dbname   = "postgres"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	CreatedAt   time.Time
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func databaseConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connection established successfully")

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := databaseConn()
	defer db.Close()

	var products []Product

	result, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		product := Product{}
		err := result.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}

	fmt.Println(products)

	templates.ExecuteTemplate(w, "Index", products)
}

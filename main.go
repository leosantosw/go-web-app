package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

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

func databaseConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disabled",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := databaseConn()
	db.Close()

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

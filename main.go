package main

import (
	"net/http"

	"web-app/routes"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":3000", nil)
}

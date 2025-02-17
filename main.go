package main

import (
	"fmt"
	"net/http"

	"web-app/routes"
)

func main() {
	routes.Routes()
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", nil)
}

package routes

import (
	"net/http"

	"web-app/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
}

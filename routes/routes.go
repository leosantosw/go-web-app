package routes

import (
	"net/http"

	"web-app/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/remove/", controllers.Remove)
	http.HandleFunc("/edit/", controllers.Edit)
}

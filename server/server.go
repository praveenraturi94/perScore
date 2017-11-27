package server

import (
	"fmt"
	"net/http"
	"perScore/app/routes"
)

// Startserver ...
func Startserver() {
	fmt.Println("server started")
	routes.Router()
	http.ListenAndServe(":8080", nil)
}

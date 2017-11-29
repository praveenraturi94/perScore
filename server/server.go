package server

import (
	"fmt"
	"net/http"
	"os"
	"perScore/app/routes"
)

// Startserver ...
func Startserver() {
	fmt.Println("server started")
	routes.Router()
	http.ListenAndServe(os.Getenv("PORT"), nil)
}

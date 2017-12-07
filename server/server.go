package server

import (
	"fmt"
	"net/http"
	"os"
	"perScore/app/routes"
)

// Startserver ...
func Startserver() {
	fmt.Println("perScore server started")
	handler := routes.Router()
	http.ListenAndServe(os.Getenv("PORT"), handler)
}

package server

import (
	"fmt"
	"net/http"
	"os"
	"perScoreServer/app/routes"
)

// Startserver ...
func Startserver() {
	fmt.Println("perScore server started on", os.Getenv("PORT"))
	handler := routes.Router()
	http.ListenAndServe(os.Getenv("PORT"), handler)
}

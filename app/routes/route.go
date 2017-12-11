package routes

import (
	"net/http"
	"perScoreServer/app/controller"

	"github.com/rs/cors"
)

// Router ...
func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/register", controller.Register)
	mux.HandleFunc("/create_question", controller.CreateQuestion)
	mux.HandleFunc("/get_question", controller.GetQuestion)
	mux.HandleFunc("/approve_entries", controller.ApproveEntries)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	return cors.Default().Handler(mux)
}

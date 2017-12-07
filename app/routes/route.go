package routes

import (
	"net/http"
	"perScore/app/controller"

	"github.com/rs/cors"
)

// Router ...
func Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", controller.Login)
	mux.HandleFunc("/register", controller.Register)
	mux.HandleFunc("/create_question", controller.CreateQues)
	mux.HandleFunc("/get_question", controller.GetQues)
	mux.HandleFunc("/approve_entries", controller.ApproveEntries)
	mux.HandleFunc("/get_entries", controller.GetEntries)
	mux.HandleFunc("/get_interests", controller.GetInterests)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	return cors.Default().Handler(mux)
}

package routes

import (
	"net/http"
	"perScore/app/controller"
)

// Router ...
func Router() {
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/create_question", controller.CreateQues)
	http.HandleFunc("/get_question", controller.GetQues)
	http.HandleFunc("/approve_entries", controller.ApproveEntries)
	http.HandleFunc("/get_entries", controller.GetEntries)
	http.HandleFunc("/get_interests", controller.GetInterests)
}

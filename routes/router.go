package routes

import (
	"github.com/gorilla/mux"

	"github.com/mathesukkj/gopassword-validator/controllers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/validate-password", controllers.ValidatePasswordHandler).Methods("post")
	return r
}

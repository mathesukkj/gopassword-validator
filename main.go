package main

import (
	"net/http"

	"github.com/mathesukkj/gopassword-validator/routes"
)

func main() {
	router := routes.NewRouter()
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

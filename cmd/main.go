package main

import (
	"log"
	"net/http"

	"github.com/mbrunoon/go-unit-converter/routes"
)

func main() {
	router := routes.NewRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}

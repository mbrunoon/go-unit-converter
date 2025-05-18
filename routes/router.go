package routes

import (
	"github.com/gorilla/mux"
	"github.com/mbrunoon/go-unit-converter/internal/converter"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", converter.HomeHandler).Methods("GET")
	router.HandleFunc("/converter", converter.ConverterHandler).Methods("POST")

	return router
}

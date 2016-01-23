package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/simplyianm/timehub/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	http.ListenAndServe(":8000", r)
}

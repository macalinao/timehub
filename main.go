package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/simplyianm/timehub/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	go func() {
		http.ListenAndServe(":8000", r)
	}()
	panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("static/"))))
}

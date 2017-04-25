package main

import (
	"log"
	"net/http"

	"github.com/prygan/Go-Rest-API/handlers"
)

func main() {
	http.HandleFunc("/outings", handlers.GetOutings)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

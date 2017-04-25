package main

import (
	"github.com/kevin/Go-Rest-API/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/outings", handlers.HandlerOutings)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

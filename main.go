package main

import (
	"fmt"
	"log"
	"net/http"
)

func getOutings(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Good Job to get Outings")
}

func main() {
	http.HandleFunc("/Outings", getOutings)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error 404 not Found, don't hack plz")
	}
}

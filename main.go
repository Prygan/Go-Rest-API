package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/satori/go.uuid"
)

type outing struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Adress string `json:adress`
	City   string `json:city`
}

func main() {
	http.HandleFunc("/Outings", getOutings)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newOuting(name string, adress string, city string) *outing {
	return &outing{
		ID:     uuid.NewV4().String(),
		Name:   name,
		Adress: adress,
		City:   city,
	}
}

func getOutings(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	outing := newOuting("sortie ouf", "qq part", "nantes")

	o, _ := json.Marshal(outing)

	fmt.Fprintf(w, string(o))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error 404 not Found, don't hack plz")
	}
}

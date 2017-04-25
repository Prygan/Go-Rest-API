package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

type outing struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
}

func newOuting(name string, address string, city string) *outing {
	return &outing{
		ID:      uuid.NewV4().String(),
		Name:    name,
		Address: address,
		City:    city,
	}
}

func HandlerOutings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		outing := newOuting("sortie ouf", "qq part", "nantes")
		o, _ := json.Marshal(outing)
		fmt.Fprintf(w, string(o))
		return
	} else if r.Method == "POST" {
		fmt.Println(r.Body)
		return
	}
	errorHandler(w, r, http.StatusNotFound)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error 404 not Found, don't hack plz")
	}
}

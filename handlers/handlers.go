package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prygan/Go-Rest-API/utils"
)

func HandlerOutings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		outings := utils.ReadJSON("data/outings.json")

		o, _ := json.Marshal(outings)

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

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prygan/Go-Rest-API/utils"
)

func GetOutings(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	outings := utils.ReadJSON("data/outings.json")

	o, _ := json.Marshal(outings)

	fmt.Fprintf(w, string(o))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Error 404 not Found, don't hack plz")
	}
}

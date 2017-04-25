package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/prygan/Go-Rest-API/models"
)

func ReadJSON(filename string) []models.Outing {
	var outings []models.Outing
	b, _ := ioutil.ReadFile(filename)
	json.Unmarshal(b, &outings)
	return outings
}

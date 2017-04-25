package models

import (
	"github.com/satori/go.uuid"
)

type Outing struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
}

func NewOuting(name string, address string, city string) *Outing {
	return &Outing{
		ID:      uuid.NewV4().String(),
		Name:    name,
		Address: address,
		City:    city,
	}
}

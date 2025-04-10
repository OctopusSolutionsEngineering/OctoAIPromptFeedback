package model

type Health struct {
	ID     string `jsonapi:"primary,health" json:"id"`
	Status string `json:"status"`
}

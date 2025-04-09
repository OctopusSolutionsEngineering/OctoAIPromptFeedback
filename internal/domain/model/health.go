package model

type Health struct {
	ID     string `jsonapi:"primary,feedback" json:"id"`
	Status string `json:"status"`
}

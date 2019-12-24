package models

// Movie model
type Movie struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}

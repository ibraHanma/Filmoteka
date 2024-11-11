package model

import "time"

type Movie struct {
	ID          int       `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	ReleaseDate time.Time `json:"ReleaseDate"`
	Rating      int       `json:"Rating"`
}

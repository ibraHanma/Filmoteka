package model

import (
	"time"
)

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title,omitempty" binding:"required"`
	Description string    `json:"description" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Rating      int       `json:"rating" binding:"required,min=1,max=10"`
}

package model

import "time"

type Actor struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Gender   string    `json:"gender"`
}

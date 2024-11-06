package model

import "time"

type Actor struct {
	ID       int
	Name     string
	Birthday time.Time
	Gender   string
}

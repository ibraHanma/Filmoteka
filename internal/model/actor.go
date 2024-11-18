package model

import (
	"time"
)

type Actor struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Birthday time.Time `json:"birthday" validate:"required"`
	Gender   string    `json:"gender" validate:"oneof=male female other"`
}

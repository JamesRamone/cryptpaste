package main

import (
	uuid "github.com/gobuffalo/uuid"
)

type Pad struct {
	ID   uuid.UUID `json:"id"`
	Data string    `json:"data"`
}

func NewPad() *Pad {
	return &Pad{
		ID:   uuid.Must(uuid.NewV4()),
		Data: "",
	}
}

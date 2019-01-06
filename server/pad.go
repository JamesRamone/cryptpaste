package main

import (
	uuid "github.com/gobuffalo/uuid"
)

type pad struct {
	ID   uuid.UUID `json:"id"`
	Data string    `json:"data"`
}

func NewPad() *pad {
	return &pad{
		ID:   uuid.Must(uuid.NewV4()),
		Data: "",
	}
}

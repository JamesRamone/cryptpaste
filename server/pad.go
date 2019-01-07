package main

import (
	"crypto/rand"

	"github.com/oklog/ulid"
)

type Pad struct {
	ID   ulid.ULID `json:"id"`
	Data string    `json:"data"`
}

func NewPad() *Pad {
	return &Pad{
		ID:   ulid.MustNew(ulid.Now(), rand.Reader),
		Data: "",
	}
}

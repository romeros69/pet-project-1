package entity

import "github.com/gofrs/uuid"

type Book struct {
	ID     uuid.UUID `json:"id"`
	Tittle string    `json:"tittle"`
	Author string    `json:"author"`
}

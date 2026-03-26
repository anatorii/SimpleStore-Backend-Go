package models

import "github.com/google/uuid"

type Image struct {
	Id          uuid.UUID `db:"id"`
	Data        []byte    `db:"data"`
	Description string    `db:"description"`
}

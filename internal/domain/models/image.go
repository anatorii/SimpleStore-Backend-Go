package models

import "github.com/google/uuid"

type Image struct {
	Id          uuid.UUID
	Data        []byte
	Description string
}

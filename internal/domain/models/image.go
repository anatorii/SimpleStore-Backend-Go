package models

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	Id          uuid.UUID `db:"id"`
	Data        []byte    `db:"data"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

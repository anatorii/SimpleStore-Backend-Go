package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	Id        uuid.UUID `db:"id"`
	Country   string    `db:"country"`
	City      string    `db:"city"`
	Street    string    `db:"street"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

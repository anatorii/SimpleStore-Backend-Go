package models

import "github.com/google/uuid"

type Address struct {
	Id      uuid.UUID `db:"id"`
	Country string    `db:"country"`
	City    string    `db:"city"`
	Street  string    `db:"street"`
}

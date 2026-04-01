package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id               uuid.UUID `db:"id"`
	ClientName       string    `db:"client_name"`
	ClientSurname    string    `db:"client_surname"`
	Birthday         time.Time `db:"birthday"`
	Gender           string    `db:"gender"`
	RegistrationDate time.Time `db:"registration_date"`
	AddressId        uuid.UUID `db:"address_id"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

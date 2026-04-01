package models

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	AddressId   uuid.UUID `db:"address_id"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

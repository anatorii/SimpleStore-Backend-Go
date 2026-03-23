package models

import "github.com/google/uuid"

type Supplier struct {
	Id          uuid.UUID
	Name        string
	AddressId   uuid.UUID
	PhoneNumber string
}

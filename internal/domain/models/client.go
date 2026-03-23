package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	Id               uuid.UUID
	ClientName       string
	ClientSurname    string
	Birthday         time.Time
	Gender           string
	RegistrationDate time.Time
	AddressId        uuid.UUID
}

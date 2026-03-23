package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id             uuid.UUID
	Name           string
	Category       string
	Price          float32
	AvailableStock int
	LastUpdateDate time.Time
	SupplierId     uuid.UUID
	ImageId        uuid.UUID
}

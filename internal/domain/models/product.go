package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	Category       string    `db:"category"`
	Price          float64   `db:"price"`
	AvailableStock int       `db:"available_stock"`
	LastUpdateDate time.Time `db:"last_update_date"`
	SupplierId     uuid.UUID `db:"supplier_id"`
	ImageId        uuid.UUID `db:"image_id"`
}

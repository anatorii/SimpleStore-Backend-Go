package dto

import (
	"storeapi/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name           string    `json:"name" validate:"required,min=3,max=255"`
	Category       string    `json:"category" validate:"required,min=3,max=255"`
	Price          float32   `json:"price" validate:"required,float"`
	AvailableStock int       `json:"available_stock" validate:"required,integer,min=1,max=1000000000"`
	LastUpdateDate time.Time `json:"last_update_date" validate:"omitempty,date"`
	SupplierId     time.Time `json:"supplier_id" validate:"omitempty,uuid"`
	ImageId        time.Time `json:"image_id" validate:"omitempty,uuid"`
}

type UpdateProductAvailableRequest struct {
	AvailableStock int `json:"available_stock" validate:"required,integer,min=1,max=1000000000"`
}

type ProductResponse struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Price          float32   `json:"price"`
	AvailableStock int       `json:"available_stock"`
	LastUpdateDate time.Time `json:"last_update_date"`
	SupplierId     time.Time `json:"supplier_id"`
	ImageId        time.Time `json:"image_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ModelToProductResponse(m *models.Product) *ProductResponse {

}

func ModelToProductResponseList(m []*models.Client) []*ProductResponse {

}

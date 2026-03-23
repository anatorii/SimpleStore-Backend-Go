package dto

import (
	"storeapi/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateSupplierRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	AddressId   string `json:"address_id" validate:"omitempty,uuid"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=255"`
}

type UpdateSupplierAddressRequest struct {
	AddressId string `json:"address_id" validate:"required,uuid"`
}

type SupplierResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	AddressId   string    `json:"address_id"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ModelToSupplierResponse(m *models.Supplier) *SupplierResponse {

}

func ModelToSupplierResponseList(m []*models.Supplier) []*SupplierResponse {

}

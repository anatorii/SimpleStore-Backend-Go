package dto

import (
	"storeapi/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateSupplierRequest struct {
	Name        string    `json:"name" validate:"required,min=3,max=255"`
	AddressId   uuid.UUID `json:"address_id" validate:"omitempty,uuid"`
	PhoneNumber string    `json:"phone_number" validate:"required,min=3,max=255"`
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
	r := SupplierResponse{
		Id:          m.Id,
		Name:        m.Name,
		PhoneNumber: m.PhoneNumber,
		AddressId:   m.AddressId.String(),
	}
	return &r
}

func ModelToSupplierResponseList(m []*models.Supplier) []*SupplierResponse {
	l := make([]*SupplierResponse, 0)
	for _, v := range m {
		r := SupplierResponse{
			Id:          v.Id,
			Name:        v.Name,
			PhoneNumber: v.PhoneNumber,
			AddressId:   v.AddressId.String(),
		}
		l = append(l, &r)
	}
	return l
}

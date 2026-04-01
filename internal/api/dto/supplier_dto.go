package dto

import (
	"fmt"
	"storeapi/internal/domain/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateSupplierRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=255" example:"Dell"`
	AddressId   string `json:"address_id" validate:"omitempty" example:"00000000-0000-0000-0000-000000000000"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=255" example:"00 1004 443"`
}

type SupplierResponse struct {
	Id          uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name        string    `json:"name" example:"Dell"`
	AddressId   string    `json:"address_id" example:"00000000-0000-0000-0000-000000000000"`
	PhoneNumber string    `json:"phone_number" example:"00 1004 443"`
	CreatedAt   time.Time `json:"created_at" example:"0001-01-01T00:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"0001-01-01T00:00:00Z"`
}

func ModelToSupplierResponse(m *models.Supplier) *SupplierResponse {
	r := SupplierResponse{
		Id:          m.Id,
		Name:        m.Name,
		PhoneNumber: m.PhoneNumber,
		AddressId:   m.AddressId.String(),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
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
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		l = append(l, &r)
	}
	return l
}

func (r *CreateSupplierRequest) Validate(validate *validator.Validate) error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	if len(r.AddressId) != 0 {
		if _, err := uuid.Parse(r.AddressId); err != nil {
			return fmt.Errorf("Invalid Address Id")
		}
	}

	return nil
}

func (r *CreateSupplierRequest) GetAddressId() uuid.UUID {
	v, _ := uuid.Parse(r.AddressId)
	return v
}

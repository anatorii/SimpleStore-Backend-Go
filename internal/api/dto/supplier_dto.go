package dto

import (
	"fmt"
	"storeapi/internal/domain/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateSupplierRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	AddressId   string `json:"address_id" validate:"omitempty"`
	PhoneNumber string `json:"phone_number" validate:"required,min=3,max=255"`
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

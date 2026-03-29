package dto

import (
	"fmt"
	"storeapi/internal/domain/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name           string  `json:"name" validate:"required,min=3,max=255"`
	Category       string  `json:"category" validate:"required,min=3,max=255"`
	Price          float64 `json:"price" validate:"gt=0"`
	AvailableStock int     `json:"available_stock" validate:"gte=0"`
	LastUpdateDate string  `json:"last_update_date" validate:"required,datetime=2006-01-02"`
	SupplierId     string  `json:"supplier_id" validate:"omitempty"`
	ImageId        string  `json:"image_id" validate:"omitempty"`
}

type UpdateProductAvailableRequest struct {
	AvailableStock int `json:"available_stock" validate:"gte=0"`
}

type ProductResponse struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Price          float64   `json:"price"`
	AvailableStock int       `json:"available_stock"`
	LastUpdateDate string    `json:"last_update_date"`
	SupplierId     uuid.UUID `json:"supplier_id"`
	ImageId        uuid.UUID `json:"image_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ModelToProductResponse(m *models.Product) *ProductResponse {
	r := ProductResponse{
		Id:             m.Id,
		Name:           m.Name,
		Price:          m.Price,
		AvailableStock: m.AvailableStock,
		LastUpdateDate: m.LastUpdateDate.Format("2006-01-02"),
		SupplierId:     m.SupplierId,
		ImageId:        m.ImageId,
	}
	return &r
}

func ModelToProductResponseList(lm []*models.Product) []*ProductResponse {
	l := make([]*ProductResponse, 0)
	for _, m := range lm {
		r := ProductResponse{
			Id:             m.Id,
			Name:           m.Name,
			Price:          m.Price,
			AvailableStock: m.AvailableStock,
			LastUpdateDate: m.LastUpdateDate.Format("2006-01-02"),
			SupplierId:     m.SupplierId,
			ImageId:        m.ImageId,
		}
		l = append(l, &r)
	}
	return l
}

func (r *CreateProductRequest) Validate(validate *validator.Validate) error {
	if err := validate.Struct(r); err != nil {
		return fmt.Errorf("Invalid request payload")
	}

	if len(r.SupplierId) != 0 {
		if _, err := uuid.Parse(r.SupplierId); err != nil {
			return fmt.Errorf("Invalid Supplier Id")
		}
	}

	if len(r.ImageId) != 0 {
		if _, err := uuid.Parse(r.ImageId); err != nil {
			return fmt.Errorf("Invalid Image Id")
		}
	}

	return nil
}

func (r *CreateProductRequest) GetLastUpdateDate() time.Time {
	t, _ := time.Parse("2006-01-02", r.LastUpdateDate)
	return t
}

func (r *CreateProductRequest) GetSupplierId() uuid.UUID {
	v, _ := uuid.Parse(r.SupplierId)
	return v
}

func (r *CreateProductRequest) GetImageId() uuid.UUID {
	v, _ := uuid.Parse(r.ImageId)
	return v
}

package dto

import (
	"fmt"
	"storeapi/internal/domain/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name           string  `json:"name" validate:"required,min=3,max=255" example:"Hammer"`
	Category       string  `json:"category" validate:"required,min=3,max=255" example:"DIY"`
	Price          float64 `json:"price" validate:"gt=0" example:"100"`
	AvailableStock int     `json:"available_stock" validate:"gte=0" example:"10"`
	LastUpdateDate string  `json:"last_update_date" validate:"required,datetime=2006-01-02" example:"2006-01-02"`
	SupplierId     string  `json:"supplier_id" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
	ImageId        string  `json:"image_id" validate:"omitempty" example:""`
}

type UpdateProductAvailableRequest struct {
	AvailableStock int `json:"available_stock" validate:"required,gte=0" example:"10"`
}

type ProductResponse struct {
	Id             uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name           string    `json:"name" example:"Hammer"`
	Category       string    `json:"category" example:"DIY"`
	Price          float64   `json:"price" example:"100"`
	AvailableStock int       `json:"available_stock" example:"10"`
	LastUpdateDate string    `json:"last_update_date" example:"2006-01-02"`
	SupplierId     uuid.UUID `json:"supplier_id" example:"00000000-0000-0000-0000-000000000000"`
	ImageId        uuid.UUID `json:"image_id" example:"00000000-0000-0000-0000-000000000000"`
	CreatedAt      time.Time `json:"created_at" example:"0001-01-01T00:00:00Z"`
	UpdatedAt      time.Time `json:"updated_at" example:"0001-01-01T00:00:00Z"`
}

func ModelToProductResponse(m *models.Product) *ProductResponse {
	r := ProductResponse{
		Id:             m.Id,
		Name:           m.Name,
		Category:       m.Category,
		Price:          m.Price,
		AvailableStock: m.AvailableStock,
		LastUpdateDate: m.LastUpdateDate.Format("2006-01-02"),
		SupplierId:     m.SupplierId,
		ImageId:        m.ImageId,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
	return &r
}

func ModelToProductResponseList(lm []*models.Product) []*ProductResponse {
	l := make([]*ProductResponse, 0)
	for _, m := range lm {
		r := ProductResponse{
			Id:             m.Id,
			Name:           m.Name,
			Category:       m.Category,
			Price:          m.Price,
			AvailableStock: m.AvailableStock,
			LastUpdateDate: m.LastUpdateDate.Format("2006-01-02"),
			SupplierId:     m.SupplierId,
			ImageId:        m.ImageId,
			CreatedAt:      m.CreatedAt,
			UpdatedAt:      m.UpdatedAt,
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

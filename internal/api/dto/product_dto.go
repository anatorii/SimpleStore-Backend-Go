package dto

import (
	"storeapi/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name           string    `json:"name" validate:"required,min=3,max=255"`
	Category       string    `json:"category" validate:"required,min=3,max=255"`
	Price          float64   `json:"price" validate:"required,gt=0"`
	AvailableStock int       `json:"available_stock" validate:"required,gte=0"`
	LastUpdateDate string    `json:"last_update_date" validate:"required,datetime=2006-01-02"`
	SupplierId     uuid.UUID `json:"supplier_id" validate:"omitempty"`
	ImageId        uuid.UUID `json:"image_id" validate:"omitempty"`
}

type UpdateProductAvailableRequest struct {
	AvailableStock int `json:"available_stock" validate:"required,gte=0"`
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

func (r *CreateProductRequest) GetLastUpdateDate() time.Time {
	t, _ := time.Parse("2006-01-02", r.LastUpdateDate)
	return t
}

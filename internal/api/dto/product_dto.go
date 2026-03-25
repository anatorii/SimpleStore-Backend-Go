package dto

import (
	"storeapi/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateProductRequest struct {
	Name           string    `json:"name" validate:"required,min=3,max=255"`
	Category       string    `json:"category" validate:"required,min=3,max=255"`
	Price          float32   `json:"price" validate:"required,float,gt=0"`
	AvailableStock int       `json:"available_stock" validate:"required,integer,gte=0"`
	LastUpdateDate time.Time `json:"last_update_date" validate:"omitempty,date"`
	SupplierId     uuid.UUID `json:"supplier_id" validate:"omitempty,uuid"`
	ImageId        uuid.UUID `json:"image_id" validate:"omitempty,uuid"`
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
		LastUpdateDate: m.LastUpdateDate,
		SupplierId:     m.SupplierId,
		ImageId:        m.ImageId,
	}
	return &r
}

func ModelToProductResponseList(m []*models.Product) []*ProductResponse {
	l := make([]*ProductResponse, 0)
	for _, v := range m {
		r := ProductResponse{
			Id:             v.Id,
			Name:           v.Name,
			Price:          v.Price,
			AvailableStock: v.AvailableStock,
			LastUpdateDate: v.LastUpdateDate,
			SupplierId:     v.SupplierId,
			ImageId:        v.ImageId,
		}
		l = append(l, &r)
	}
	return l
}

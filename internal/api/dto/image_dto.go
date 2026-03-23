package dto

import (
	"storeapi/internal/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateImageRequest struct {
	Image       []byte `json:"image" validate:"required"`
	Description string `json:"name" validate:"required,min=3,max=1024"`
}

type UpdateImageRequest struct {
	Image       []byte `json:"image" validate:"required"`
	Description string `json:"name" validate:"required,min=3,max=1024"`
}

type ImageResponse struct {
	Id          uuid.UUID `json:"id"`
	Image       []byte    `json:"image" validate:"required"`
	Description string    `json:"name" validate:"required,min=3,max=1024"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ModelToImageResponse(m *models.Image) *ImageResponse {

}

package dto

import (
	"encoding/base64"
	"storeapi/internal/domain/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateImageRequest struct {
	ProductId uuid.UUID `json:"product_id" validate:"required"`
	Data      string    `json:"image" validate:"required,base64"`
}

type UpdateImageRequest struct {
	Data string `json:"image" validate:"required,base64"`
}

type ImageResponse struct {
	Id        uuid.UUID `json:"id"`
	Data      string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ModelToImageResponse(m *models.Image) *ImageResponse {
	r := &ImageResponse{
		Id:   m.Id,
		Data: base64.StdEncoding.EncodeToString(m.Data),
	}
	return r
}

func (r *CreateImageRequest) Validate(validate *validator.Validate) error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	return nil
}

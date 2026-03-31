package dto

import (
	"encoding/base64"
	"fmt"
	"storeapi/internal/domain/models"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateImageRequest struct {
	ProductId uuid.UUID `json:"product_id" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
	Data      string    `json:"image" validate:"required" example:"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="`
}

type UpdateImageRequest struct {
	Data string `json:"image" validate:"required" example:"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="`
}

type ImageResponse struct {
	Id        uuid.UUID `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Data      string    `json:"image" example:"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="`
	CreatedAt time.Time `json:"created_at" example:"0001-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"0001-01-01T00:00:00Z"`
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

	_, err := base64.StdEncoding.DecodeString(r.Data)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}

	return nil
}

func (r *UpdateImageRequest) Validate(validate *validator.Validate) error {
	if err := validate.Struct(r); err != nil {
		return err
	}

	_, err := base64.StdEncoding.DecodeString(r.Data)
	if err != nil {
		return fmt.Errorf("failed to decode base64: %w", err)
	}

	return nil
}

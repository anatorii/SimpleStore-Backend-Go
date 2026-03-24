package repos

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type ImageRepoInt interface {
	GetById(ctx context.Context, id uuid.UUID) (*models.Image, error)
	GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error)
	Create(ctx context.Context, model *models.Image) error
	Update(ctx context.Context, model *models.Image) error
	Delete(ctx context.Context, id uuid.UUID) error
}

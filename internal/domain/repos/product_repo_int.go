package repos

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type ProductRepoInt interface {
	GetAll(ctx context.Context) ([]*models.Product, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Product, error)
	Create(ctx context.Context, model *models.Product) error
	Update(ctx context.Context, model *models.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

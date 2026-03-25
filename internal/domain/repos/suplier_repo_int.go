package repos

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type SupplierRepoInt interface {
	GetAll(ctx context.Context) ([]*models.Supplier, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Supplier, error)
	Create(ctx context.Context, model *models.Supplier) error
	Update(ctx context.Context, model *models.Supplier) error
	Delete(ctx context.Context, id uuid.UUID) error
}

package repos

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type ClientRepoInt interface {
	GetAll(ctx context.Context) ([]*models.Client, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Client, error)
	GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error)
	Create(ctx context.Context, model *models.Client) error
	Update(ctx context.Context, model *models.Client) error
	Delete(ctx context.Context, id uuid.UUID) error
}

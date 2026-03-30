package service

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type ClientService interface {
	GetAll(ctx context.Context, offset, limit int) ([]*models.Client, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Client, error)
	GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error)
	Create(ctx context.Context, model *models.Client) error
	Update(ctx context.Context, model *models.Client) error
	UpdateAddress(ctx context.Context, model *models.Client, address models.Address) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type SupplierService interface {
	GetAll(ctx context.Context) ([]*models.Supplier, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Supplier, error)
	Create(ctx context.Context, model *models.Supplier) error
	Update(ctx context.Context, model *models.Supplier) error
	UpdateAddress(ctx context.Context, model *models.Supplier, address models.Address) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductService interface {
	GetAll(ctx context.Context) ([]*models.Product, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Product, error)
	Create(ctx context.Context, model *models.Product) error
	Update(ctx context.Context, model *models.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ImageService interface {
	GetById(ctx context.Context, id uuid.UUID) (*models.Image, error)
	GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error)
	Create(ctx context.Context, model *models.Image, product *models.Product) error
	Update(ctx context.Context, model *models.Image) error
	Delete(ctx context.Context, id uuid.UUID) error
}

package repository

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type AddressRepo interface {
	GetById(ctx context.Context, id uuid.UUID) (*models.Address, error)
	GetByAddress(ctx context.Context, country, city, street string) (*models.Address, error)
	Create(ctx context.Context, model *models.Address) error
	Update(ctx context.Context, model *models.Address) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ClientRepo interface {
	GetAll(ctx context.Context) ([]*models.Client, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Client, error)
	GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error)
	Create(ctx context.Context, model *models.Client) error
	Update(ctx context.Context, model *models.Client) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type SupplierRepo interface {
	GetAll(ctx context.Context) ([]*models.Supplier, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Supplier, error)
	Create(ctx context.Context, model *models.Supplier) error
	Update(ctx context.Context, model *models.Supplier) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ProductRepo interface {
	GetAll(ctx context.Context) ([]*models.Product, error)
	GetById(ctx context.Context, id uuid.UUID) (*models.Product, error)
	Create(ctx context.Context, model *models.Product) error
	Update(ctx context.Context, model *models.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ImageRepo interface {
	GetById(ctx context.Context, id uuid.UUID) (*models.Image, error)
	GetByProductId(ctx context.Context, productId uuid.UUID) (*models.Image, error)
	Create(ctx context.Context, model *models.Image) error
	Update(ctx context.Context, model *models.Image) error
	Delete(ctx context.Context, id uuid.UUID) error
}

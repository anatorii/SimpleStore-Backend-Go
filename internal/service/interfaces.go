package service

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
)

type ClientService interface {
	GetById(ctx context.Context, id uuid.UUID) (*models.Client, error)
	GetByName(ctx context.Context, fullname models.FullName) (*models.Client, error)
}

type SupplierService interface {
}

type ProductService interface {
}

type ImageService interface {
}

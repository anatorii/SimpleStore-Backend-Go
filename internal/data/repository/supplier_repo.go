package repository

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SupplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) *SupplierRepo {
	return &SupplierRepo{db: db}
}

func (r *SupplierRepo) GetAll(ctx context.Context) ([]*models.Supplier, error) {
	return nil, nil
}

func (r *SupplierRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Supplier, error) {
	return nil, nil
}

func (r *SupplierRepo) Create(ctx context.Context, model *models.Supplier) error {
	return nil
}

func (r *SupplierRepo) Update(ctx context.Context, model *models.Supplier) error {
	return nil
}

func (r *SupplierRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

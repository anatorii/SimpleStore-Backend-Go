package repository

import (
	"context"
	"storeapi/internal/domain/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) GetAll(ctx context.Context) ([]*models.Product, error) {
	return nil, nil
}

func (r *ProductRepo) GetById(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return nil, nil
}

func (r *ProductRepo) Create(ctx context.Context, model *models.Product) error {
	return nil
}

func (r *ProductRepo) Update(ctx context.Context, model *models.Product) error {
	return nil
}

func (r *ProductRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
